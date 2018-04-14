package config

import (
	cfg"github.com/Unknwon/goconfig"
	"DailyServer/commons/util"
)

var Cfg *cfg.ConfigFile

const CONFIG string = "/config.ini"

func InitConfig() (err error) {
	rootDir := util.GetCurrentDirectory()
	Cfg, err = cfg.LoadConfigFile(rootDir + CONFIG)
	return err
}
