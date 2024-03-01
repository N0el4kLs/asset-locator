package runner

import "fmt"

const VERSION = "v0.1.0"

func ShowBanner() {
	//http://www.network-science.de/ascii/  smslant
	var banner = `
   __                 __          
  / /  ___  _______ _/ /____  ____
 / /__/ _ \/ __/ _ '/ __/ _ \/ __/
/____/\___/\__/\_,_/\__/\___/_/ %s
`
	fmt.Printf(banner, VERSION)
}
