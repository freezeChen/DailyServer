package mysql

import (
	"fmt"
	. "DailyServer/commons/config"
	. "DailyServer/commons/log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

const CONN_CAP int = 2048

var (
	err    error
	engine *xorm.Engine
	once   sync.Once
)

func Refresh() {
	once.Do(func() {
		db_name := Cfg.MustValue("mysql", "db_name", "")
		connstring := Cfg.MustValue("mysql", "connstring", "")
		fmt.Println(db_name,connstring)
		engine, err = xorm.NewEngine("mysql", connstring+"/"+db_name+"?charset=utf8")
		if err != nil {
			LogFile.I("mysql", fmt.Sprintf("Failed to newengine: %s", err))
		} else {
			engine.ShowSQL(true)
			engine.TZLocation = time.Local
			engine.SetMaxOpenConns(CONN_CAP)
			engine.SetMaxIdleConns(CONN_CAP)
			engine.SetMapper(core.SameMapper{})
		}
	})
	if err = engine.Ping(); err != nil {
		LogFile.I("mysql", fmt.Sprintf("Failed to ping: %s", err))
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
