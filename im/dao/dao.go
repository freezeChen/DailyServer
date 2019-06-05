
/*
   @Time : 2019-05-31 11:39:47
   @Author :
   @File : dao
   @Software: im
*/
package dao

import (
	"dailyserver/im/conf"
	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Dao struct {
	Db    xorm.EngineInterface
	Redis *redis.Redis
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		Db:    mysql.New(c.Mysql),
		Redis: redis.New(c.Redis),
	}
	return
}

