package main

import "github.com/N0el4kLs/asset-locator/lib"

func main() {
	target := []string{"https://www.baidu.com", "jd.com"}
	al := lib.NewAssetLocatorEngine(lib.DefaultResultCallback).
		UnSetWeight().
		LoadTargets(target)
	al.RunSearch()
}
