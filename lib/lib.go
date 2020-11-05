package lib

import (
	"fmt"
)

func Run(ncuUser *CASUser, watchTime int) {
	iptvWatcher := Login(ncuUser.Username, ncuUser.Password)
	iptvWatcher.GetBaseVideoUrl()
	//iptvWatcher.watchTime = watchTime
	fmt.Println("baseURL:", iptvWatcher.baseUrl)
	iptvWatcher.StartRequest()
	fmt.Println("start watching video")
	times := watchTime * 3
	for i := 0; i < times; i++ {
		iptvWatcher.KeepWatchRequest()
	}
	iptvWatcher.EndRequest()
	fmt.Println("end watching video")
	wg.Done()
}