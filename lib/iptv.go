package lib

import (
	"fmt"
	"io/ioutil"
	"iptv-helper/util"
	"net/http"
	"time"
)

type Iptv struct {
	iptvUsername string
	iptvPassword string
	baseUrl string
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

func (instance *Iptv)GetBaseVideoUrl() {
	fmt.Println("userVideoUrl",instance.userVideoUrl())
	resp, err := http.Get(instance.userVideoUrl())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	respPlain, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("respPlain",string(respPlain))
	instance.baseUrl = util.ParseXmlToUrl(&respPlain)
}