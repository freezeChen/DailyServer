/*
   @Time : 2019-06-17 16:46:43
   @Author : 
   @File : dao
   @Software: job
*/
package dao

import (
	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"dailyserver/job/conf"
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