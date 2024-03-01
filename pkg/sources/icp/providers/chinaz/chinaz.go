package chinaz

import (
	"errors"
	"fmt"
	"strings"

	"asset-locator/pkg/sources/icp/providers"
	"asset-locator/pkg/util"

	"github.com/PuerkitoBio/goquery"
)

var (
	CHINAZ     = "ICPChinaz"
	CHINAZ_URL = "https://seo.chinaz.com/%s"

	ChinazErr   = errors.New("chinaz icp search error")
	ICPNotFound = errors.New("icp not found")
)

type ICPEngine struct {
}

func (I ICPEngine) Name() string {
	return CHINAZ
}

func (I ICPEngine) SearchICP(domain string) (string, error) {
	u := fmt.Sprintf(CHINAZ_URL, domain)
	client := util.DefaultClient
	header := map[string]string{
		"Accept":             "*/*",
		"Accept-Language":    "zh-CN,zh;q=0.9,en;q=0.8",
		"Cookie":             "cz_statistics_visitor=d41a1fd6-9ca7-6fac-6356-968008591e6e; Hm_lvt_aecc9715b0f5d5f7f34fba48a3c511d6=1709270204; _clck=q3u99v%7C2%7Cfjp%7C0%7C1489; qHistory=aHR0cDovL3Nlby5jaGluYXouY29tX1NFT+e7vOWQiOafpeivog==; Hm_lvt_ca96c3507ee04e182fb6d097cb2a1a4c=1708448385,1708669675,1709270219; toolbox_urls=baidu.com|39.156.66.10; Hm_lpvt_aecc9715b0f5d5f7f34fba48a3c511d6=1709270282; _clsk=l6k63x%7C1709270815649%7C5%7C1%7Ct.clarity.ms%2Fcollect; Hm_lpvt_ca96c3507ee04e182fb6d097cb2a1a4c=1709271217",
		"Origin":             "https://seo.chinaz.com",
		"Referer":            "https://seo.chinaz.com/",
		"Sec-Ch-Ua":          `"Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"`,
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": "macOS",
		"Sec-Fetch-Dest":     "empty",
		"Sec-Fetch-Mode":     "cors",
		"Sec-Fetch-Site":     "same-site",
		"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
	}
	resp, err := client.Get(u, header)
	if err != nil {
		return providers.None, ChinazErr
	}
	html, _ := resp.ToString()
	rst := getICPInfo(html)
	rst.URL = domain
	if rst.IcpName == "" {
		return providers.None, ICPNotFound
	}
	return rst.Hostingparty, nil
}

func getICPInfo(html string) result {
	var rst result
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return rst
	}
	hostingpartySelector := "#company > i > a"
	typeSelector := "body > div._chinaz-seo-new1.wrapper.mb10 > table > tbody > tr:nth-child(4) > td._chinaz-seo-newtc._chinaz-seo-newh40 > span:nth-child(3) > i"
	icpSelector := "body > div._chinaz-seo-new1.wrapper.mb10 > table > tbody > tr:nth-child(4) > td._chinaz-seo-newtc._chinaz-seo-newh40 > span:nth-child(1) > i > a"
	hostingparty := doSelect(hostingpartySelector, doc)
	icpType := doSelect(typeSelector, doc)
	icpName := doSelect(icpSelector, doc)

	rst.Hostingparty = hostingparty
	rst.IcpType = icpType
	rst.IcpName = icpName

	return rst
}

func doSelect(selector string, doc *goquery.Document) string {
	var rst string
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		rst = s.Text()
	})
	return rst
}

type result struct {
	URL          string
	Hostingparty string
	IcpName      string
	IcpType      string
}
