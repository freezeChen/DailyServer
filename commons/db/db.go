package db

import (
	. "dailyserver2/commons/config"
	"dailyserver2/commons/glog"
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
	engine *xorm.EngineGroup
	once   sync.Once
)

func InitDb() {
	once.Do(func() {
		db_name := Cfg.MustValue("mysql", "db_name", "")
		master := Cfg.MustValue("mysql", "master", "")
		slave := Cfg.MustValue("mysql", "slave", "")
		conns := []string{
			master + "/" + db_name + "?charset=utf8",
			slave + "/" + db_name + "?charset=utf8",
		}
		engine, err = xorm.NewEngineGroup("mysql", conns, xorm.LeastConnPolicy())
		if err != nil {
			glog.Info("mysql", fmt.Sprintf("Failed to newengine: %s", err))
		} else {
			engine.ShowSQL(true)
			engine.TZLocation = time.Local
			engine.SetMaxOpenConns(CONN_CAP)
			engine.SetMaxIdleConns(CONN_CAP)
			engine.SetMapper(core.SnakeMapper{})
		}
	})
	if err = engine.Ping(); err != nil {
		glog.Info("mysql", fmt.Sprintf("Failed to ping: %s", err))
	}
}

func NewEngine() (*xorm.EngineGroup, error) {
	return engine, err
}

func Ping() error {
	return engine.Ping()
}

func Close() {
	engine.Close()
}
