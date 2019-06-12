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
	"dailyserver/im/server"
)

type Service struct {
	dao *dao.Dao
	srv *server.Server
}

func New(c *conf.Config, server *server.Server) (s *Service) {
	s = &Service{
		dao: dao.New(c),
		srv: server,
	}
	return s
}
