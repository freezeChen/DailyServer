package config

import (
	"DailyServer/commons/util"
	cfg "github.com/Unknwon/goconfig"
)

const CONFIG string = "/config.ini"

var Cfg *cfg.ConfigFile

var DefaultConfig *Config

type Config struct {
	Domain       string
	AbsolutePath string
	VirtualPath  string
	RedisConn    string
	Redispwd     string
}

func SetConf() (err error) {
	rootDir := util.GetCurrentDirectory()
	Cfg, err = cfg.LoadConfigFile(rootDir + CONFIG)

	if err == nil {
		DefaultConfig, err = InitConfig()
	}

	return err
}

func InitConfig() (cfg *Config, err error) {
	cfg = new(Config)
	cfg.RedisConn = Cfg.MustValue("redis", "conn", "")
	cfg.Redispwd = Cfg.MustValue("redis", "pwd", "")

	return cfg, err
}
