/*
   @Time : 2019-06-17 16:46:43
   @Author :
   @File : main
   @Software: job
*/
package main

import (
	"dailyserver/job/conf"
	"dailyserver/job/service"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/kafka"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}
	svc := micro.NewService(
		micro.Name("go.micro.srv.job"),
		micro.Address(":8085"),

	)
	svc.Init()
	//imService := proto.NewIMService("go.micro.srv.im", svc.Client())
	micro.Broker(kafka.NewBroker(func(options *broker.Options) {
		options.Addrs = []string{"www.frozens.vip:9092"}
	}))

	s := service.New(cfg)

	//err = proto.RegisterHelloHandler(svc.Server(), s)
	if err := svc.Run(); err != nil {
		return
	}
}
