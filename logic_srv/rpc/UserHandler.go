/*
   @Time : 2018/8/31 下午2:13
   @Author :
   @File : LogicHandler
   @Software: DailyServer
*/
package rpc

import (
	"DailyServer/commons/glog"
	"DailyServer/commons/gredis"
	"DailyServer/commons/util"
	"DailyServer/grpc"
	"DailyServer/logic_srv/models"
	"context"
)

type LogicHandler struct{}

func (LogicHandler) Receive(ctx context.Context, req *grpc.ReceiveReq, reply *grpc.ReceiveReply) error {

	switch req.Proto.Opr {
	case grpc.OpSendMsg:
		var msg models.Message
		msg.Uid = util.ToInt64(req.Proto.Id)
		msg.Rid = util.ToInt64(req.Proto.Toid)
		msg.Msg = string(req.Proto.Body)
		msg.Type = 1

		err := models.InsertMsg(&msg)
		if err != nil {
			glog.Infof("Failed to insert msg:%s", err)
			return err
		}

	}

	return nil
}

func (LogicHandler) Check(ctx context.Context, req *grpc.CheckReq, reply *grpc.CheckReply) error {
	user, err := models.GetUserByID(int64(req.Id))
	if err != nil {
		return err
	}
	gredis.Set()
	reply.Key = int32(user.Id)
	return nil
}
