package util

import (
	"fmt"

	"github.com/joeguo/tldextract"
)

func ExtractDomain(u string) string {
	extracted, _ := tldextract.New("tld.cache", false)
	rst := extracted.Extract(u)
	//fmt.Printf("%+v;%s\n", rst, u)
	if rst.Flag == tldextract.Domain {
		return fmt.Sprintf("%s.%s", rst.Root, rst.Tld)
	}
	//fmt.Println("Can't extract domain: ", u)
	return u
}
