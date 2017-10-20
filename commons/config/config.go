package config

import (
	"github.com/Unknwon/goconfig"
	"DailySever/commons/util"
)

var Cfg goconfig.ConfigFile

const CONFIG  string ="\\config.ini"

func InitConfig() (err error){
	rootDir := util.GetCurrentDirectory()
	Cfg, err :=goconfig.LoadConfigFile(rootDir+CONFIG)
	return err
}