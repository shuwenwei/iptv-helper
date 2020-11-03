package main

import (
	"iptv-helper/lib"
	"iptv-helper/util"
)

func main() {
	iptvConfig := lib.LoadCfg(util.DefaultCfgPath)
	//ncuUser := iptvConfig.NcuUser
	//appConfig := iptvConfig.AppConfig
	factory := lib.IptvFactory{iptvConfig}
	factory.CreateTasks()
	//lib.Run(&ncuUser, &appConfig)
}
