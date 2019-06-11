/*
   @Time : 2019-06-10 13:42:13
   @Author :
   @File : service
   @Software: logic
*/
package service

import (
	"dailyserver/logic/conf"
	"dailyserver/logic/dao"
)

type Service struct {
	dao *dao.Dao
}

func New(c *conf.Config) (s *Service) {
	s = &Service{
		dao: dao.New(c),
	}

	return s
}
