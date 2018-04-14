package models

import (
	"github.com/go-xorm/xorm"
	"DailyServer/commons/mysql"
	"DailyServer/commons/log"
)

func Engine() *xorm.Engine {
	Engine, err := mysql.NewEngine()
	if err != nil {
		log.LogFile.E("failed to newengine")
	}
	return Engine
}
