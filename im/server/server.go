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
	"errors"
	"fmt"
	"time"
)

const (
	_HandshakeTimeout = 8
	_HeartBeat        = 10 * time.Minute
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

func (server *Server) AuthTCP(ctx context.Context, msg *proto.Proto, ch *model.Channel) (id int64, err error) {

	if err := msg.ReadTCP(&ch.Reader); err != nil {
		return 0, err
	}

	if msg.Opr != proto.OpAuth {
		err = errors.New("authTCP op is error")
	}

	//string(msg.Body)

	return
}

func (server *Server) Heartbeat(ctx context.Context, id int64) (err error) {
	fmt.Println("Heartbeat")
	return
}

func (server *Server) Operate(ctx context.Context, msg *proto.Proto, ch *model.Channel) (err error) {
	fmt.Println("Operate")
	return
}

func (server *Server) DisConnect(ctx context.Context, id int64) (err error) {
	fmt.Println("DisConnect")
	return nil
}
