/*
   @Time : 2019-01-31 10:35
   @Author : frozenchen
   @File : main
   @Software: DailyServer
*/
package main

import (
	handler2 "DailyServer/back/cmd/job/handler"
	glog2 "DailyServer/back/commons/glog"
	"DailyServer/back/constant"
	grpc2 "DailyServer/back/grpc"
	"github.com/micro/cli"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/kafka"
	_ "github.com/micro/go-plugins/broker/kafka"
)

func main() {
	microService := micro.NewService(
		micro.Name(constant.MICRO_JOB_SRV),
		micro.RegisterTTL(constant.MICRO_TTL),
		micro.RegisterInterval(constant.MICRO_Interval),
	)
	microService.Init(
		micro.Action(
			func(context *cli.Context) {
				glog2.InitLogger()

			},
		), )

	micro.Broker(kafka.NewBroker(func(options *broker.Options) {
		options.Addrs = []string{
			"www.frozens.vip:9092",
		}
	}))
	broker.Init()

	if err := broker.Connect(); err != nil {
		glog2.Painc(err)
	}

	imService := grpc2.NewIMService(constant.MICRO_IM_SRV, microService.Client())

	handler2.NewJobHandler(imService).Start()

	if err := microService.Run(); err != nil {
		glog2.Painc(err)
	}

}
