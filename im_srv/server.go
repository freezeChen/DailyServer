/*
   @Time : 2019-01-30 15:11
   @Author : frozenchen
   @File : Server
   @Software: DailyServer
*/
package im

import (
	"DailyServer/grpc"
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
	return s.bucket.Get(key)
}

func (s *Server) Online(proto *grpc.Proto, ch *Channel) {
	s.bucket.Online(proto.Id, ch)
}

func (srv *Server) Connect(ctx context.Context, id int32) (int32,error) {
	reply, err := srv.logicService.Check(ctx, &grpc.CheckReq{Id: id})
	return reply.Key,err
}
