/*
   @Time : 2019-06-17 16:46:43
   @Author : 
   @File : service
   @Software: job
*/
package service

import (
	"dailyserver/job/conf"
	"dailyserver/job/dao"
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

