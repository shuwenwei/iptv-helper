package lib

import (
	"fmt"
	"time"
)

func Run(ncuUser *CASUser, appConfig *AppConfig) {
	iptvWatcher := Login(ncuUser.Username, ncuUser.Password)
	iptvWatcher.cfg = appConfig
	iptvWatcher.GetBaseVideoUrl()
	fmt.Println("baseURL:", iptvWatcher.baseUrl)
	go iptvWatcher.StartRequest()
	fmt.Println("start watching video")
	time.Sleep(time.Minute)
	iptvWatcher.EndRequest()
}