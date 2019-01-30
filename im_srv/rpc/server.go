/*
   @Time : 2019-01-30 16:31
   @Author : frozenchen
   @File : server
   @Software: DailyServer
*/
package rpc

import (
	"DailyServer/im_srv"
	"DailyServer/grpc"
	"context"
)

type Server struct {
	srv *im.Server
}

func NewImServer(srv *im.Server) *Server {
	var server = Server{}
	server.srv = srv
	return &server
}

func (s *Server) PushMsg(ctx context.Context, req *grpc.PushMsgReq, reply *grpc.PushMsgReply) error {
	if channel := s.srv.Channel(req.Key); channel != nil {

		channel.Push(req.Proto)
	}
}
