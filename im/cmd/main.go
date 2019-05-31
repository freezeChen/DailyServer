/*
   @Time : 2019-05-31 11:39:47
   @Author :
   @File : main
   @Software: im
*/
package main

import (
	"dailyserver/im/conf"
	"dailyserver/im/server"
	"dailyserver/im/service"
	"dailyserver/proto"
	"github.com/micro/go-micro"
	"time"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}
	svc := micro.NewService(
		micro.Name("go.micro.srv.hello"),
		micro.Address(":8081"),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(20*time.Second),
	)
	svc.Init()

	s := service.New(cfg)

	if err := server.InitTCP(s, cfg); err != nil {
		//log.error
		return
	}

	err = proto.RegisterHelloHandler(svc.Server(), s)
	if err := svc.Run(); err != nil {
		return
	}
}
