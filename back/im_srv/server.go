/*
   @Time : 2019-01-30 15:11
   @Author : frozenchen
   @File : Server
   @Software: DailyServer
*/
package im

import (
	"DailyServer/back/grpc"
	"context"
)

type Server struct {
	logicService grpc.LogicService
	bucket       *Bucket
}

func NewServer(logicService grpc.LogicService) *Server {
	var server = &Server{logicService: logicService}
	server.bucket = NewBucket()
	return server
}

func (s *Server) Channel(key int32) *Channel {
	return Get(key)
}

func (s *Server) Online(proto *grpc.Proto, ch *Channel) {
	Online(proto.Id, ch)
}

func (srv *Server) Connect(ctx context.Context, id int32) (int32, error) {
	reply, err := srv.logicService.Check(ctx, &grpc.CheckReq{Id: id})
	if err != nil {
		return 0,err
	}
	return reply.Key, err
}

//处理接收的消息
func (srv *Server) Operate(ctx context.Context, proto *grpc.Proto, ch *Channel) error {
	_, err := srv.logicService.Receive(ctx, &grpc.ReceiveReq{Proto: proto})
	return err
}
