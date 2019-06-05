/*
   @Time : 2019-05-31 15:23
   @Author : frozenchen
   @File : server
   @Software: DailyServer
*/
package server

import (
	"context"
	"dailyserver/im/model"
	"dailyserver/im/service"
	"dailyserver/proto"
)

type Server struct {
	svc *service.Service
}

func New(svc *service.Service) *Server {
	return &Server{svc: svc}
}

func (s *Server) AuthTCP(ctx context.Context, proto *proto.Proto, ch *model.Channel) (err error) {
	return
}