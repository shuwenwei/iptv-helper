package lib

import (
	"fmt"
	"time"
)

func Run(ncuUser *CASUser) {
	iptvWatcher := Login(ncuUser.Username, ncuUser.Password)
	iptvWatcher.GetBaseVideoUrl()
	fmt.Println("baseURL:", iptvWatcher.baseUrl)
	go iptvWatcher.StartRequest()
	fmt.Println("start watching video")
	time.Sleep(time.Minute)
	iptvWatcher.EndRequest()
	fmt.Println("end watching video")
	wg.Done()
}