package pearktrue

import (
	"errors"
	"fmt"

	"asset-locator/pkg/sources/icp/providers"
	"asset-locator/pkg/util"
)

// https://api.pearktrue.cn/api/icp/?domain=baidu.com

var (
	PEARKTRUE = "Pearktrue"
	ICP_URL   = "https://api.pearktrue.cn/api/icp/?domain=%s"

	PearktrueErr           = errors.New("pearktrue icp search error")
	PearktrueUnmarkshalErr = errors.New("pearktrue icp unmarshal error")
	ICPNotFound            = errors.New("icp not found")
)

type ICPEngine struct {
}

func (i ICPEngine) Name() string {
	return PEARKTRUE
}
func (i ICPEngine) SearchICP(domain string) (string, error) {
	u := fmt.Sprintf(ICP_URL, domain)
	client := util.DefaultClient
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Appl" +
			"eWebKit/537.36 (KHTML, like Gecko) Chrome/",
	}
	resp, err := client.Get(u, headers)
	if err != nil {
		return providers.None, PearktrueErr
	}
	var rst result

	if err = resp.UnmarshalJson(&rst); err != nil {
		return providers.None, PearktrueUnmarkshalErr
	}
	if hostingPart := rst.Data.Hostingparty; hostingPart != "" {
		return hostingPart, nil
	} else {
		return providers.None, ICPNotFound
	}
}

type result struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Domain string `json:"domain"`
	Data   struct {
		Hostingparty string `json:"hostingparty"`
		Filingnumber string `json:"filingnumber"`
		Websitename  string `json:"websitename"`
		Audittime    string `json:"audittime"`
	} `json:"data"`
	ApiSource string `json:"api_source"`
}
