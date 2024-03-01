package webscan

import (
	"encoding/json"
	"fmt"

	"asset-locator/pkg/util"

	"github.com/projectdiscovery/gologger"
)

// https://www.webscan.cc/api/index.html
var (
	WEBSCAN_URL = "http://api.webscan.cc/?action=query&ip=%s"
)

type WebScanEngine struct {
}

func (w WebScanEngine) Ip2Domain(ip string) (string, error) {
	u := fmt.Sprintf(WEBSCAN_URL, ip)
	client := util.DefaultClient
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Appl" +
			"eWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		"Host":   "api.webscan.cc",
		"Pragma": "no-cache",
	}
	resp, err := client.Get(u, headers)
	if err != nil {
		gologger.Debug().Msgf("Webscan ip to domain error:,url: %s\n", u)
		return "", err
	}
	body, err := resp.ToString()
	fmt.Println(body)
	var rst []result
	err = json.Unmarshal(resp.Bytes(), &rst)
	if err != nil {
		gologger.Debug().Msgf("Webscan unmarshall result error:,url: %s\n", u)
		return "", err
	}
	fmt.Println(rst)
	return "", nil
}

type result struct {
	Domain string `json:"domain"`
	Title  string `json:"title"`
}
