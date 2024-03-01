package ip38

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"asset-locator/pkg/util"

	"github.com/projectdiscovery/gologger"
)

var (
	None     = ""
	IP38     = "Ip38"
	IP38_URL = "https://site.ip138.com/%s/"

	Ip38Err = errors.New("ip38 ip to domain error")
)

type Ip38Engine struct {
}

func (i Ip38Engine) Name() string {
	return IP38
}

func (i Ip38Engine) SearchIP2Domain(ip string) (string, error) {
	u := fmt.Sprintf(IP38_URL, ip)
	client := util.DefaultClient
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Appl" +
			"eWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp," +
			"image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Connection":                "close",
		"sec-ch-ua":                 "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"98\", \"Google Chrome\";v=\"98\"",
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        "\"Windows\"",
		"Upgrade-Insecure-Requests": "1",
		"Sec-Fetch-Site":            "same-origin",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-User":            "?1",
		"Sec-Fetch-Dest":            "document",
		"Referer":                   "https://site.ip138.com/",
		"Accept-Language":           "zh-CN,zh;q=0.9",
	}
	resp, err := client.Get(u, headers)
	if err != nil {
		return None, Ip38Err
	}
	sBody, err := resp.ToString()
	if err != nil {
		return None, Ip38Err
	}
	if strings.Contains(sBody, "绑定过的域名如下：") && !strings.Contains(sBody, "<li>暂无结果</li>") {
		pattern := regexp.MustCompile(`<li class="title"><span>绑定过的域名如下：</span></li>((?s).*)</ul>`)
		match := pattern.FindString(sBody)
		domainPattern := regexp.MustCompile(`<a.*?>(.*?)</a>`)
		domains := domainPattern.FindAllStringSubmatch(match, -1)
		for _, d := range domains {
			fmt.Println(d[1])
		}
	}
	return None, nil
}

func (i Ip38Engine) Ip2Domain(ip string) {
	u := fmt.Sprintf(IP38_URL, ip)
	client := util.DefaultClient
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Appl" +
			"eWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp," +
			"image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"Connection":                "close",
		"sec-ch-ua":                 "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"98\", \"Google Chrome\";v=\"98\"",
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        "\"Windows\"",
		"Upgrade-Insecure-Requests": "1",
		"Sec-Fetch-Site":            "same-origin",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-User":            "?1",
		"Sec-Fetch-Dest":            "document",
		"Referer":                   "https://site.ip138.com/",
		"Accept-Language":           "zh-CN,zh;q=0.9",
	}
	resp, err := client.Get(u, headers)
	if err != nil {
		gologger.Debug().Msgf("Ip38 ip to domain error:,url: %s\n", u)
	}
	stringBody := resp.String()
	if strings.Contains(stringBody, "绑定过的域名如下：") && !strings.Contains(stringBody, "<li>暂无结果</li>") {
		pattern := regexp.MustCompile(`<li class="title"><span>绑定过的域名如下：</span></li>((?s).*)</ul>`)
		match := pattern.FindString(stringBody)
		domainPattern := regexp.MustCompile(`<a.*?>(.*?)</a>`)
		domains := domainPattern.FindAllStringSubmatch(match, -1)
		//fmt.Println(domains)
		for _, d := range domains {
			fmt.Println(d[1])
		}
	}
}
