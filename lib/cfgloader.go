package lib

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"iptv-helper/util"
	"os"
)

type IptvConfig struct {
	NcuUser NcuUser		`toml:"user"`
}

type NcuUser struct {
	Username string
	Password string
}

func LoadCfg(path string) *NcuUser {
	var iptvConfig IptvConfig
	_, err := toml.DecodeFile(path,&iptvConfig)
	if err != nil {
		panic(err)
	}
	return &iptvConfig.NcuUser
}

func setCache(videoUserPwd string) {
	f, err := os.Open(util.DefaultCachePath)
	if err != nil {
		fmt.Println(err)
		f, err = os.Create(util.DefaultCachePath)
		if err != nil {
			return
		}
	}
	_, err = f.WriteString(videoUserPwd)
	if err != nil {
		fmt.Println(err)
	}
}

func checkCache() (string, bool) {
	f, err := os.Open(util.DefaultCachePath)
	if err != nil {
		return "", false
	}
	body, _ := ioutil.ReadAll(f)
	return string(body), true
}