package main

import (
	"iptv-helper/lib"
	"iptv-helper/util"
)

func main() {
	iptvConfig := lib.LoadCfg(util.DefaultCfgPath)
	factory := lib.IptvFactory{iptvConfig}
	factory.CreateTasks()
}
