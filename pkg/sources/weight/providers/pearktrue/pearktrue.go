package pearktrue

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/N0el4kLs/asset-locator/pkg/sources/weight/providers"
	"github.com/N0el4kLs/asset-locator/pkg/util"
)

// https://api.pearktrue.cn/api/website/weight.php?domain=www.baidu.com

var (
	WEIGHT_URL = "https://api.pearktrue.cn/api/website/weight.php?domain=%s"

	PearktrueError           = errors.New("pearktrue weight engine error")
	PearktrueUnmarshallError = errors.New("pearktrue weight engine unmarshal error")
)

type WeightEngine struct {
}

func (w WeightEngine) Name() string {
	return "Pearktrue"
}

func (w WeightEngine) SearchWeight(domain string) (providers.WeightLevel, error) {
	u := fmt.Sprintf(WEIGHT_URL, domain)
	client := util.DefaultClient
	resp, err := client.Get(u, nil)
	if err != nil {
		return providers.ErrorLevel, PearktrueError
	}
	var rst result
	if err := resp.UnmarshalJson(&rst); err != nil {
		return providers.ErrorLevel, PearktrueUnmarshallError
	}

	// get weight
	if rst.Code == 200 && rst.Data.BaiDuPC != "" {
		weight, _ := strconv.Atoi(rst.Data.BaiDuPC)
		if weight == 0 {
			return providers.LowLevel, nil
		} else if weight > 0 {
			return providers.CriticalLevel, nil
		} else {
			return providers.ErrorLevel, nil
		}
	} else {
		return providers.ErrorLevel, PearktrueError
	}
}

type result struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Domain string `json:"domain"`
	Data   struct {
		BaiDuPC string `json:"BaiDu_PC"`
	} `json:"data"`
}
