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
	"github.com/micro/go-micro"
	"time"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}
	svc := micro.NewService(
		micro.Name("go.micro.srv.logic"),
		micro.Address(":8070"),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(20*time.Second))

	svc.Init()
	s := service.New(cfg)
	err = proto.RegisterLogicHandler(svc.Server(), s)
	if err := svc.Run(); err != nil {
		return
	}
}
