/*
   @Time : 2019-06-10 13:42:13
   @Author :
   @File : main
   @Software: logic
*/
package main

import (
	"dailyserver/logic/conf"
	"dailyserver/logic/service"
	"dailyserver/proto"
	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/kafka"
	"time"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}
	zlog.InitLogger(cfg.Log)
	svc := micro.NewService(
		micro.Name("go.micro.srv.logic"),
		micro.Address(":8070"),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(20*time.Second))

	svc.Init()

	micro.Broker(kafka.NewBroker(func(options *broker.Options) {
		options.Addrs = []string{"47.106.137.3:9092"}
	}))

	if err := broker.Connect(); err != nil {
		panic(err)
		return
	}


	s := service.New(cfg)
	err = proto.RegisterLogicHandler(svc.Server(), s)
	if err := svc.Run(); err != nil {
		return
	}
}
