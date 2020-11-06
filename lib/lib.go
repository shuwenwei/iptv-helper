package lib

import (
	"fmt"
)

func Run(iptvWatcher *Iptv) {
	baseUrl := iptvWatcher.getBaseUrl()
	defer func() {
		iptvWatcher.EndRequest(baseUrl)
		fmt.Println("end watching video")
		wg.Done()
	}()
	iptvWatcher.StartRequest(baseUrl)
	fmt.Println("start watching video")
	times := iptvWatcher.watchTime * 3
	for i := 0; i < times && !iptvWatcher.endFlag; i++ {
		iptvWatcher.KeepWatchRequest(baseUrl)
	}
}