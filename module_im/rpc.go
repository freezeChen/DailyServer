package main

import (
	"context"
	"dailyserver2/commons/glog"
	"dailyserver2/proto"
	"github.com/micro/go-micro"
)

type IMHandler struct {
}

func (IMHandler) Ping(c context.Context, req *proto.PingReq, res *proto.PingRes) error {
	glog.Info(req.Id)
	return nil
}

func (IMHandler) PushMsg(c context.Context, req *proto.PushMsgReq, res *proto.PushMsgRes) error {
	glog.Info(req.Id)
	return nil
}

func (IMHandler) PushMsgs(c context.Context, req *proto.PushMsgsReq, res *proto.PushMsgsRes) error {
	glog.Info(req.Id)
	return nil
}

func InitRPC() (err error) {

	go func() {
		service := micro.NewService(micro.Name("module_im"))
		err = proto.RegisterIMServiceHandler(service.Server(), new(IMHandler))
		if err != nil {
			glog.Errorf("Failed to start imRPC server:%s", err)
			return
		}
		err = service.Run()
		if err != nil {
			glog.Errorf("Failed to run imRPC server:%s", err)
		}
		return
	}()

	return
}
