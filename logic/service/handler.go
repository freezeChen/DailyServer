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
	"dailyserver/logic/model"
	"dailyserver/proto"
	"github.com/freezeChen/studio-library/zlog"
	"time"
)

func (s *Service) Auth(ctx context.Context, req *proto.AuthReq, reply *proto.EmptyReply) error {

	zlog.Error("auth")
	user := s.dao.GetUser(req.Id)

	if user == nil {
		return zerrors.NewMsg("user is empty")
	}

	return nil
}

func (s *Service) Operate(ctx context.Context, req *proto.OperateReq, reply *proto.EmptyReply) error {
	switch req.Proto.Opr {
	case proto.OpSendMsg:
		msg := &model.SingleInfo{
			Uid:  req.Proto.Id,
			ToId: req.Proto.ToId,
			Msg:  string(req.Proto.Body),
			Time: time.Now().String(),
		}

		if err := s.dao.InsertSingleInfo(msg); err != nil {
			return err
		}

		if err := s.dao.Kafka.SendSingleMsg(msg); err != nil {
			return err
		}
	}

	return nil
}
