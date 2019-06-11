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
)

func (s *Service) Auth(ctx context.Context, req *proto.AuthReq, reply *proto.EmptyReply) error {

	user := s.dao.GetUser(req.Id)

	if user == nil {
		return zerrors.NewMsg("user is empty")
	}

	return nil
}
