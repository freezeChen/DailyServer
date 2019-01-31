/*
   @Time : 2018/8/31 下午2:13
   @Author :
   @File : LogicHandler
   @Software: DailyServer
*/
package rpc

import (
	"DailyServer/grpc"
	"DailyServer/logic_srv/models"
	"context"
)

type LogicHandler struct{}

func (LogicHandler) Receive(ctx context.Context, req *grpc.ReceiveReq, reply *grpc.ReceiveReply) error {

	panic("implement me")
}

func (LogicHandler) Check(ctx context.Context, req *grpc.CheckReq, reply *grpc.CheckReply) error {
	user, err := models.GetUserByID(int64(req.Id))
	if err != nil {
		return err
	}
	reply.Key = int32(user.Id + 1)
	return nil
}
