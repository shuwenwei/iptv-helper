package lib

import (
	"context"
	"fmt"
	"io/ioutil"
	"iptv-helper/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type IptvFactory struct {
	IptvCfg *IptvConfig
}

var wg sync.WaitGroup

func (factory *IptvFactory) CreateTasks() {
	cfg := factory.IptvCfg
	tasknum := cfg.AppConfig.Tasknum
	if tasknum <= 0 {
		log.Fatal("illegal tasknum:", tasknum)
	}
	iptvWatcher := Login(cfg.NcuUser.Username, cfg.NcuUser.Password)
	iptvWatcher.watchTime = cfg.AppConfig.Tasknum
	iptvWatcher.endFlag = false
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for s := range signalChannel{
			switch s {
			case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM:
				iptvWatcher.endFlag = true
				fmt.Println("exit in 20 seconds")
			}
		}
	}()
	for i := 0; i < tasknum; i++ {
		wg.Add(1)
		go Run(iptvWatcher)
	}
	wg.Wait()
}

type Iptv struct {
	iptvUsername string
	iptvPassword string
	watchTime int
	endFlag bool
	//baseUrl string
}

func (instance *Iptv) userVideoUrl() string {
	testUrl := fmt.Sprint("http://wyjx.ncu.edu.cn/VIEWGOOD/adi/portal/load.ashx?" +
		"ModeType=PlayVOD," +
		"StreamType=HTTP_MP4," +
		"Ver=8.0.0.2," +
		"StreamID=7037," +
		"ClassID=53," +
		"ClassName=%e5%bd%b1%e8%a7%86%e6%ac%a3%e8%b5%8f," +
		"assetID=193," +
		"assetName=%e7%88%86%e8%a3%82%e9%bc%93%e6%89%8b," +
		"Episode_ID=1," +
		"Username="+ instance.iptvUsername + "," +
		"Password=" + instance.iptvPassword + "," +
		"Redirect=false," +
		"Random="+ fmt.Sprintf("%v", time.Now().Unix()) + "000")
	return testUrl
}

func (instance *Iptv) getBaseUrl() string {
	fmt.Println("userVideoUrl",instance.userVideoUrl())
	resp, err := http.Get(instance.userVideoUrl())
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	respPlain, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("respPlain",string(respPlain))

	return util.ParseXmlToUrl(&respPlain)
}

func (instance *Iptv)StartRequest(ctx context.Context) {
	videoStartUrl := fmt.Sprintf("%s%sRandom=%v000", ctx.Value("baseUrl"), util.VideoStartSuffix, time.Now().Unix())
	fmt.Println(videoStartUrl)
	resp, err :=  http.Get(videoStartUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
}

func (instance *Iptv) KeepWatchRequest(ctx context.Context) {
	videoRunningUrl := fmt.Sprintf("%s%sRandom=%v000", ctx.Value("baseUrl"), util.VideoRunningSuffix, time.Now().Unix())
	resp, err := http.Get(videoRunningUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(time.Second * 20)
	defer resp.Body.Close()
}

func (instance *Iptv) EndRequest(ctx context.Context) {
	videoEndUrl := fmt.Sprintf("%s%s", ctx.Value("baseUrl"), util.VideoTeardownSuffix)
	fmt.Println(videoEndUrl)
	resp, _ := http.Get(videoEndUrl)
	defer resp.Body.Close()
}