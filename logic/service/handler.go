/*
   @Time : 2019-06-10 13:42:13
   @Author :
   @File : service
   @Software: logic
*/
package service

import (
	"context"
	"dailyserver/lib/zerrors"
	"dailyserver/proto"
	"github.com/freezeChen/studio-library/zlog"
)

func (s *Service) Auth(ctx context.Context, req *proto.AuthReq, reply *proto.EmptyReply) error {

	zlog.Error("auth")
	user := s.dao.GetUser(req.Id)

	if user == nil {
		return zerrors.NewMsg("user is empty")
	}

	return nil
}
