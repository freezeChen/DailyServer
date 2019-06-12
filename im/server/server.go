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
	Round  *model.Round
	Bucket *model.Bucket
	logic  proto.LogicService
}

func New(logic proto.LogicService) *Server {
	return &Server{
		Round:  model.NewRound(),
		Bucket: model.NewBucket(),
		logic:  logic,
	}
}

//验证账号
func (server *Server) AuthTCP(ctx context.Context, msg *proto.Proto, ch *model.Channel) (id int64, err error) {
	if err := msg.ReadTCP(&ch.Reader); err != nil {
		return 0, err
	}

	if msg.Opr != proto.OpAuth {
		err = errors.New("authTCP op is error")
		return
	}

	_, err = server.logic.Auth(ctx, &proto.AuthReq{
		Id: msg.Id,
	})

	if err != nil {
		return
	}

	id = msg.Id

	//string(msg.Body)

	return
}

func (server *Server) Heartbeat(ctx context.Context, id int64) (err error) {
	fmt.Println("Heartbeat")
	return
}

func (server *Server) Operate(ctx context.Context, msg *proto.Proto, ch *model.Channel) (err error) {

	switch msg.Opr {
	case proto.OpSendMsg:
		//server
	}

	return
}

func (server *Server) DisConnect(ctx context.Context, id int64) (err error) {
	fmt.Println("DisConnect")
	return nil
}
