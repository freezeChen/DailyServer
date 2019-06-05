
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
	"github.com/freezeChen/studio-library/zlog"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"time"
)

func main() {
	cfg, err := conf.Init()
	if err != nil {
		panic("load config error:" + err.Error())
	}

	zlog.InitLogger(cfg.Log)

	s := service.New(cfg)

	svc := micro.NewService(
		micro.Name("go.micro.srv.hello"),
		micro.Address(":8081"),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(20*time.Second),
	)
	svc.Init(micro.Action(func(context *cli.Context) {

	}))

	if err := server.InitTCP(s, cfg); err != nil {
		panic("initTCP error:" + err.Error())
		return
	}

	if err := proto.RegisterHelloHandler(svc.Server(), s); err != nil {
		panic("register hello error:" + err.Error())
		return
	}

	if err := svc.Run(); err != nil {
		panic("micro run error:" + err.Error())
		return
	}
}

