/*
   @Time : 2019-05-31 11:39:47
   @Author :
   @File : service
   @Software: im
*/
package service

import (
	"context"
	"dailyserver/proto"
	"fmt"
)

func (s *Service) Hello(ctx context.Context, req *proto.Req, reply *proto.Reply) error {

	reply.Message = fmt.Sprintf("hello %s, Congratulations you success call rpc service!", req.S)
	return nil
}
