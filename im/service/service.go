/*
   @Time : 2019-05-31 11:39:47
   @Author : 
   @File : service
   @Software: im
*/
package service

import (
	"dailyserver/im/conf"
	"dailyserver/im/dao"
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

