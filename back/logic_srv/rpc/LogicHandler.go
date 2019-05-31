/*
   @Time : 2018/8/31 下午2:13
   @Author :
   @File : LogicHandler
   @Software: DailyServer
*/
package rpc

import (
	"DailyServer/back/commons/glog"
	"DailyServer/back/commons/util"
	"DailyServer/back/grpc"
	"DailyServer/back/lib"
	cache2 "DailyServer/back/logic_srv/cache"
	models2 "DailyServer/back/logic_srv/models"
	publish2 "DailyServer/back/logic_srv/publish"
	"context"
	"encoding/json"
)

type LogicHandler struct {
	kafkaPub *publish2.KafkaPub
}

func NewLogicHandler(kafkaPub *publish2.KafkaPub) *LogicHandler {
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

		var msg models2.Message
		msg.Uid = info.Id
		msg.Rid = info.Rid
		msg.Msg = info.Msg
		msg.Type = 1

		err = models2.InsertMsg(&msg)
		if err != nil {
			glog.Infof("Failed to insert msg:%s", err)
			return err
		}

		if online := cache2.UserIsOnline(info.Rid); online {
			err = logic.kafkaPub.PushSingleMsg(util.ToString(msg.Rid), req.Proto)
			if err != nil {
				glog.Error(err)
				return err
			}
		}
	}

	return nil
}

func (srv LogicHandler) Check(ctx context.Context, req *grpc.CheckReq, reply *grpc.CheckReply) error {
	user, err := models2.GetUserByID(int64(req.Id))
	if err != nil {
		return err
	}
	reply.Key = int32(user.Id)
	err = cache2.SaveOnlineUser(reply.Key)
	if err != nil {
		glog.Error("save user", err)
		return err
	}

	err = srv.kafkaPub.AuthSuccess(util.ToString(req.Id))

	return err
}
