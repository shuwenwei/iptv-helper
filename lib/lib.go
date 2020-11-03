package lib

import (
	"fmt"
	"time"
)

func Run(username, password string) {
	iptvWatcher := Login(username, password)
	iptvWatcher.GetBaseVideoUrl()
	fmt.Println("baseURL:", iptvWatcher.baseUrl)
	go iptvWatcher.StartRequest()
	fmt.Println("start watching video")
	time.Sleep(time.Minute)
	iptvWatcher.EndRequest()
}