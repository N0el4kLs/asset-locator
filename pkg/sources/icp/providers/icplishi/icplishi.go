package icplishi

import (
	"errors"
	"fmt"
	"strings"

	"asset-locator/pkg/sources/icp/providers"
	"asset-locator/pkg/util"

	"github.com/PuerkitoBio/goquery"
)

var (
	ICP_LISHI = "ICPLishi"
	ICP_URL   = "https://icplishi.com/%s/"

	ICPLishiErr           = errors.New("icplishi icp search error")
	ICPLishiUnmarkshalErr = errors.New("icplishi icp unmarshal error")
	ICPNotFound           = errors.New("icp not found")
)

type ICPEngine struct {
}

func (I ICPEngine) Name() string {
	return ICP_LISHI
}

func (I ICPEngine) SearchICP(domain string) (string, error) {
	u := fmt.Sprintf(ICP_URL, domain)
	client := util.DefaultClient
	headers := map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Appl" +
			"eWebKit/537.36 (KHTML, like Gecko) Chrome/",
		"X-Forwarded-For": "127.0.0.1",
	}
	resp, err := client.Get(u, headers)
	if err != nil {
		return providers.None, ICPLishiErr
	}
	html, _ := resp.ToString()
	rst := getICPInfo(html)
	if rst.IcpName == "" {
		return providers.None, ICPNotFound
	}
	return rst.IcpName, nil
}

func getICPInfo(html string) result {
	var rst result
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return rst
	}
	urlSelector := "body > div.wrapper > div.container > div > div.module.mod-panel > div.bd > div:nth-child(1) > div.c-bd > table > tbody > tr:nth-child(1) > td:nth-child(2) > span"
	icpTypeSelector := "body > div.wrapper > div.container > div > div.module.mod-panel > div.bd > div:nth-child(1) > div.c-bd > table > tbody > tr:nth-child(2) > td:nth-child(2) > span"
	hostingpartySelector := "body > div.wrapper > div.container > div > div.module.mod-panel > div.bd > div:nth-child(1) > div.c-bd > table > tbody > tr:nth-child(3) > td:nth-child(2) > a"
	icpSelector := "body > div.wrapper > div.container > div > div.module.mod-panel > div.bd > div:nth-child(1) > div.c-bd > table > tbody > tr:nth-child(4) > td:nth-child(2) > a"

	url := doSelect(urlSelector, doc)
	hostingparty := doSelect(hostingpartySelector, doc)
	icpType := doSelect(icpTypeSelector, doc)
	icpName := doSelect(icpSelector, doc)

	rst.URL = url
	rst.Hostingparty = hostingparty
	rst.IcpType = icpType
	rst.IcpName = icpName
	return rst
}

func doSelect(selector string, doc *goquery.Document) string {
	var rst string
	doc.Find(selector).
		Each(func(i int, s *goquery.Selection) {
			text := s.Text()
			text = strings.TrimSpace(text)
			rst = text
		})
	return rst
}

type result struct {
	URL          string
	Hostingparty string
	IcpName      string
	IcpType      string
}
