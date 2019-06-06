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

const (
	_HandshakeTimeout = 8
)

type Server struct {
	Round *model.Round

	svc *service.Service
}

func New(svc *service.Service) *Server {
	return &Server{
		Round: model.NewRound(),
		svc:   svc,
	}
}

func (server *Server) AuthTCP(ctx context.Context, proto *proto.Proto, ch *model.Channel) (id int64, err error) {
	return
}
