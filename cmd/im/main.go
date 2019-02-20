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
	"DailyServer/im_srv/rpc"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	var server *im.Server
	//labelSelector := label.NewSelector()

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


	logicService := grpc.NewLogicService(constant.MICRO_LOGIC_SRV, microService.Client())

	server = im.NewServer(logicService)

	if err := im.InitTCP(server); err != nil {
		glog.Painc(err)
	}

	if err := grpc.RegisterIMServiceHandler(microService.Server(), rpc.NewImServer(server)); err != nil {
		glog.Painc(err)
	}

	if err := microService.Run(); err != nil {
		glog.Painc(err)
	}

}
