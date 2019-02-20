/*
   @Time : 2018/8/31 下午2:13
   @Author :
   @File : LogicHandler
   @Software: DailyServer
*/
package rpc

import (
	"DailyServer/commons/glog"
	"DailyServer/commons/util"
	"DailyServer/grpc"
	"DailyServer/lib"
	"DailyServer/logic_srv/cache"
	"DailyServer/logic_srv/models"
	"DailyServer/logic_srv/publish"
	"context"
	"encoding/json"
)

type LogicHandler struct {
	kafkaPub *publish.KafkaPub
}

func NewLogicHandler(kafkaPub *publish.KafkaPub) *LogicHandler {
	return &LogicHandler{kafkaPub: kafkaPub}
}

func (logic LogicHandler) Receive(ctx context.Context, req *grpc.ReceiveReq, reply *grpc.ReceiveReply) error {

	switch req.Proto.Opr {
	case grpc.OpSendMsg:
		var info lib.Info

		err := json.Unmarshal(req.Proto.Body, &info)
		if err != nil {
			return err
		}

		var msg models.Message
		msg.Uid = info.Id
		msg.Rid = info.Rid
		msg.Msg = info.Msg
		msg.Type = 1

		err = models.InsertMsg(&msg)
		if err != nil {
			glog.Infof("Failed to insert msg:%s", err)
			return err
		}

		if online := cache.UserIsOnline(info.Rid); online {
			err = logic.kafkaPub.PushSingleMsg(util.ToString(msg.Rid), req.Proto)
			if err != nil {
				glog.Error(err)
				return err
			}
		}

	}

	return nil
}

func (LogicHandler) Check(ctx context.Context, req *grpc.CheckReq, reply *grpc.CheckReply) error {
	user, err := models.GetUserByID(int64(req.Id))
	if err != nil {
		return err
	}
	reply.Key = int32(user.Id)
	err = cache.SaveOnlineUser(reply.Key)
	if err != nil {
		glog.Error("save user", err)
	}
	return nil
}
