package main

import (
	"fmt"
	"iptv-helper/lib"
	"iptv-helper/util"
)

func main() {
	ncuUser := lib.LoadCfg(util.DefaultCfgPath)
	fmt.Println(ncuUser.Username)
	fmt.Println(ncuUser.Password)
	lib.Run(ncuUser.Username, ncuUser.Password)
}
