package main

import (
	"DailyServer/commons/glog"
	"DailyServer/proto"
	"context"
	"github.com/micro/go-micro"
	"time"
)

type IMHandler struct {
}

func (IMHandler) Ping(c context.Context, req *proto.PingReq, res *proto.PingRes) error {
	glog.Info(string(req.Id))
	return nil
}

func (IMHandler) PushMsg(c context.Context, req *proto.PushMsgReq, res *proto.PushMsgRes) error {
	glog.Info(string(req.Id))
	return nil
}

func (IMHandler) PushMsgs(c context.Context, req *proto.PushMsgsReq, res *proto.PushMsgsRes) error {
	glog.Info(string(req.Id))
	return nil
}

func InitRPC() (err error) {

	go func() {
		service := micro.NewService(micro.Name("module_im"),
			micro.RegisterTTL(25*time.Second),
			micro.RegisterInterval(20*time.Second),
		)
		service.Init()
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
