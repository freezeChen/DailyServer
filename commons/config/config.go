package config

import (
	cfg"github.com/Unknwon/goconfig"
	"DailySever/commons/util"
	"fmt"
)

var Cfg *cfg.ConfigFile

const CONFIG string = "\\config.ini"

func InitConfig() (err error) {
	rootDir := util.GetCurrentDirectory()
	fmt.Println("fff:",rootDir)
	Cfg, err = cfg.LoadConfigFile(rootDir + CONFIG)
	return err
}
