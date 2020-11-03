package lib

import (
	"github.com/BurntSushi/toml"
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