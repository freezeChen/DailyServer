/*
   @Time : 2019-01-30 15:05
   @Author : frozenchen
   @File : im
   @Software: DailyServer
*/
package main

import (
	"DailyServer/commons/glog"
	"DailyServer/constant"
	"DailyServer/grpc"
	"DailyServer/im_srv"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/kafka"
)

func main() {
	var server *im.Server

	microService := micro.NewService(
		micro.Name(constant.MICRO_IM_SRV),
		micro.RegisterTTL(constant.MICRO_TTL),
		micro.RegisterInterval(constant.MICRO_Interval),
	)
	microService.Init(
		micro.Action(func(context *cli.Context) {
			glog.InitLogger()
			im.InitIMConfig()
		}))

	micro.Broker(kafka.NewBroker(func(options *broker.Options) {
		options.Addrs = []string{"test1:8092"}
	}))

	broker.Init()

	if err := broker.Connect(); err != nil {
		glog.Painc(err)
	}



	logicService := grpc.NewLogicService(constant.MICRO_LOGIC_SRV, microService.Client())

	server = im.NewServer(logicService)

	if err := im.InitTCP(server); err != nil {
		glog.Painc(err)
	}

	microService.Run()
	//if err := im.InitRPC(); err != nil {
	//	glog.Painc(err)
	//}

	//glog.Debug("im RPC is running...")
	////glog.Debug("logic server is running ...")
	//http.HandleFunc("/ws", im.InitWebSocket)
	//http.ListenAndServe(":8888", nil)

	im.InitSignal()
}
