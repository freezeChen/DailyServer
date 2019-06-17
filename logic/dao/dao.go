/*
   @Time : 2019-06-10 13:42:13
   @Author :
   @File : dao
   @Software: logic
*/
package dao

import (
	"dailyserver/logic/conf"
	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Dao struct {
	db    xorm.EngineInterface
	Redis *redis.Redis
	Kafka *kafka
}

func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		db:    mysql.New(c.Mysql),
		Redis: redis.New(c.Redis),
		Kafka: new(kafka),
	}
	return
}
