package db

import (
	"DailyServer/back/commons/config"
	glog2 "DailyServer/back/commons/glog"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"sync"
	"time"
)

const CONN_CAP int = 2048

var (
	err    error
	engine *xorm.Engine
	once   sync.Once
)

func InitDb() {
	once.Do(func() {
		db_name := config.Cfg.MustValue("mysql", "db_name", "")
		master := config.Cfg.MustValue("mysql", "master", "")

		source := master + "/" + db_name + "?charset=utf8"

		engine, err = xorm.NewEngine("mysql", source)
		if err != nil {
			glog2.Info("mysql", fmt.Sprintf("Failed to newengine: %s", err))
		} else {
			engine.ShowSQL(true)
			engine.TZLocation = time.Local
			engine.SetMaxOpenConns(CONN_CAP)
			engine.SetMaxIdleConns(CONN_CAP)
			engine.SetMapper(core.SnakeMapper{})
		}
	})
	if err = engine.Ping(); err != nil {
		glog2.Info("mysql", fmt.Sprintf("Failed to ping: %s", err))
	}
}

func NewEngine() (*xorm.Engine, error) {
	return engine, err
}

func Ping() error {
	return engine.Ping()
}

func Close() {
	engine.Close()
}

func Engine() *xorm.Engine {
	group, _ := NewEngine()
	return group
}
