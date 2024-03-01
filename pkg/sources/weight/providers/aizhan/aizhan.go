package aizhan

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"asset-locator/pkg/sources/weight/providers"
	"asset-locator/pkg/util"

	"github.com/projectdiscovery/gologger"
)

var (
	WEIGHT_URL = "https://baidurank.aizhan.com/baidu/%s/"
)

type WeightEngine struct {
}

func (w WeightEngine) Name() string {
	return "Aizhan"
}

// SearchWeight search weight via aizhan
func (w WeightEngine) SearchWeight(domianOrIP string) (providers.WeightLevel, error) {
	searchWeigthURL := fmt.Sprintf(WEIGHT_URL, domianOrIP)
	client := util.DefaultClient
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Appl" +
			"eWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		"Host":   "baidurank.aizhan.com",
		"Pragma": "no-cache",
	}
	resp, err := client.Get(searchWeigthURL, headers)
	stringBody := resp.String()
	if stringBody == "" {
		gologger.Debug().Msgf("Get weight return NULL body,url: %s\n", searchWeigthURL)
		return providers.ErrorLevel, err
	}
	wt, err := getWeight(stringBody)

	if err != nil {
		return providers.ErrorLevel, err
	}
	return wt, nil
}

// getWeight extract weight from html
func getWeight(s string) (providers.WeightLevel, error) {
	pattern := regexp.MustCompile(`images/br/([0-9]{1,2})\.png`)
	w := pattern.FindStringSubmatch(s)
	if len(w) == 2 {
		weightS := w[1]
		weightI, _ := strconv.Atoi(weightS)
		if providers.WeightLevel(weightI) > providers.LowLevel {
			return providers.CriticalLevel, nil
		}
		return providers.LowLevel, nil
	}
	return providers.ErrorLevel, errors.New("get weight failed")
}
