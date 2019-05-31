/*
   @Time : 2019-01-30 15:05
   @Author : frozenchen
   @File : im
   @Software: DailyServer
*/
package main

import (
	glog2 "DailyServer/back/commons/glog"
	"DailyServer/back/constant"
	grpc2 "DailyServer/back/grpc"
	"DailyServer/back/im_srv"
	rpc2 "DailyServer/back/im_srv/rpc"
	"github.com/micro/cli"
)

func main() {
	var server *im_srv.Server
	//labelSelector := label.NewSelector()

	microService := micro.NewService(
		micro.Name(constant.MICRO_IM_SRV),
		micro.RegisterTTL(constant.MICRO_TTL),
		micro.RegisterInterval(constant.MICRO_Interval),

	)
	microService.Init(
		micro.Action(func(context *cli.Context) {
			glog2.InitLogger()
			im_srv.InitIMConfig()
		}))


	logicService := grpc2.NewLogicService(constant.MICRO_LOGIC_SRV, microService.Client())

	server = im_srv.NewServer(logicService)

	if err := im_srv.InitTCP(server); err != nil {
		glog2.Painc(err)
	}

	if err := grpc2.RegisterIMServiceHandler(microService.Server(), rpc2.NewImServer(server)); err != nil {
		glog2.Painc(err)
	}

	if err := microService.Run(); err != nil {
		glog2.Painc(err)
	}

}
