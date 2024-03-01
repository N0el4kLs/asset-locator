package runner

import (
	"bufio"
	"net"
	"net/url"
	"os"
	"strings"

	"asset-locator/pkg/sources/icp"
	"asset-locator/pkg/sources/weight"
	"asset-locator/pkg/sources/weight/providers"
	"asset-locator/pkg/util"

	"github.com/projectdiscovery/gologger"
)

type TargetType int

const (
	ErrorType TargetType = iota - 1
	IP
	Domain
	URL
)

type Runner struct {
	Options *Options

	// Support formats:
	// IP address: 8.8.8.8 ;Domain: baidu.com ;URL: https://baidu.com https://127.0.0.1:8080
	Target []string

	Callback func(rst Result)
}

func NewRunner(option *Options) (*Runner, error) {
	runner := &Runner{
		Options: option,
	}

	if hasStdin() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			if isValidateTarget(s.Text()) != ErrorType {
				runner.Target = append(runner.Target, s.Text())
			}
		}
	}
	if option.Target != "" {
		runner.Target = append(runner.Target, option.Target)
	}

	if option.Targets != "" {
		if targets, err := util.ReadFileByLine(option.Targets); err != nil {
			return runner, err
		} else {
			runner.Target = append(runner.Target, targets...)
		}
	}

	gologger.Info().Msgf("Target count: %d\n", len(runner.Target))

	runner.Callback = DefaultCallout

	return runner, nil
}

func (r *Runner) Run() error {
	for _, t := range r.Target {
		tValue, tType := parseTarget(t)
		rst := Result{
			Target: t,
			Weight: providers.ErrorLevel,
		}
		if r.Options.Weight && tType == Domain {
			level, err := weight.SearchWeight(tValue.(string))
			if err != nil {
				gologger.Error().Msgf("Get Weight error: %s\n", err)
			}
			rst.Weight = level
		}
		if tType == Domain {
			if icp, err := icp.SearchICP(tValue.(string)); err != nil {
				gologger.Error().Msgf("Get ICP error: %s\n", err)
			} else {
				rst.ICP = icp
			}
		}

		r.Callback(rst)
	}
	return nil
}

func (r *Runner) Close() {

}

func hasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	return (stat.Mode() & os.ModeCharDevice) == 0
}

// check if the target is a valid target
func isValidateTarget(t string) TargetType {
	if strings.HasPrefix(t, "http") {
		if _, err := url.Parse(t); err != nil {
			return ErrorType
		} else {
			return URL
		}
	}

	// Try parsing as an IP address
	if ip := net.ParseIP(t); ip != nil {
		return IP
	}
	if strings.Count(t, ".") > 0 { // Todo this logic is not accurate
		return Domain
	}
	return ErrorType
}

func parseTarget(t string) (interface{}, TargetType) {
	tType := isValidateTarget(t)

	switch tType {
	case URL: // if input is url
		u, _ := url.Parse(t)
		return parseTarget(u.Hostname())
	case IP: // if input is ip
		return t, IP
	case Domain:
		return t, Domain
	default:
		return nil, ErrorType
	}
}
