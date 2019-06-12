/*
   @Time : 2019-05-31 11:39:47
   @Author :
   @File : service
   @Software: im
*/
package service

import (
	"context"
	"dailyserver/lib/zerrors"
	"dailyserver/proto"
)

func (s *Service) PushMsg(ctx context.Context, req *proto.PushMsgReq, reply *proto.PushMsgReply) error {
	ch := s.srv.Bucket.Get(req.Key)
	if ch == nil {
		return zerrors.NewMsg("用户不在线")
	}

	ch.Push(req.Proto)
	return nil
}
