/*
   @Time : 2019-01-30 16:31
   @Author : frozenchen
   @File : server
   @Software: DailyServer
*/
package rpc

import (
	"DailyServer/back/commons/glog"
	"DailyServer/back/grpc"
	"DailyServer/back/im_srv"
	"context"
	"github.com/micro/go-micro/errors"
)

type Server struct {
	srv *im_srv.Server
}

func NewImServer(srv *im_srv.Server) *Server {
	var server = Server{}
	server.srv = srv
	return &server
}

func (s Server) PushMsg(ctx context.Context, req *grpc.PushMsgReq, reply *grpc.PushMsgReply) error {
	glog.Debug("push msg",req)
	if channel := s.srv.Channel(req.Proto.Id); channel != nil {


		err := channel.Push(req.Proto)
		return err
	} else {
		return errors.New("", "channel is not online", 200)
	}

	return nil
}
