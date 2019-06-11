/*
   @Time : 2019-06-10 13:42:13
   @Author :
   @File : dao
   @Software: logic
*/
package dao

import (
	"dailyserver/logic/conf"
	"fmt"
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

func (d *Dao) Get() {
	type us struct {
		Id   int64
		Name string
	}

	rows, e := d.Db.SQL("select id from user;").Rows(new(string))
	if e != nil {
		panic(e)
		return
	}

	for rows.Next() {
		var u string
		err := rows.Scan(&u)
		if err != nil {
			panic(e)
		}
		fmt.Println(fmt.Sprintf("%+v", u))
	}

}
