package runner

import (
	"flag"
	"fmt"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

type Options struct {
	// target
	Target  string
	Targets string

	// scan mode
	Weight bool // weight scan

	Debug bool
}

func ParseOptions() (*Options, error) {
	ShowBanner()
	options := &Options{}

	flag.Usage = usage
	flag.StringVar(&options.Target, "t", "", "target")
	flag.StringVar(&options.Targets, "T", "", "file with targets")
	flag.BoolVar(&options.Weight, "w", false, "weight scan")
	flag.BoolVar(&options.Debug, "debug", false, "debug mode")

	flag.Parse()
	if options.Debug {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	} else {
		gologger.DefaultLogger.SetMaxLevel(levels.LevelInfo)
	}
	return options, nil
}

func usage() {
	comment := `This is a tool for Web application/IP asset ownership query, especially design for butian src
Usage of ./assetlocator:
  -T string
    	file with targets
  -t string 
		target
  -w bool	
		weight scan
  -debug bool
		debug mode
`
	fmt.Println(comment)
}
