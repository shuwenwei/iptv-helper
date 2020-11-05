package lib

import (
	"context"
	"fmt"
)

func Run(iptvWatcher *Iptv) {
	ctx := context.WithValue(context.TODO(), "baseUrl", iptvWatcher.getBaseUrl())
	defer func() {
		iptvWatcher.EndRequest(ctx)
		fmt.Println("end watching video")
		wg.Done()
	}()
	iptvWatcher.StartRequest(ctx)
	fmt.Println("start watching video")
	times := iptvWatcher.watchTime * 3
	for i := 0; i < times && !iptvWatcher.endFlag; i++ {
		iptvWatcher.KeepWatchRequest(ctx)
	}
}