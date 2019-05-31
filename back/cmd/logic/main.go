package main

import (
	config2 "DailyServer/back/commons/config"
	db2 "DailyServer/back/commons/db"
	glog2 "DailyServer/back/commons/glog"
	gredis2 "DailyServer/back/commons/gredis"
	"DailyServer/back/constant"
	grpc2 "DailyServer/back/grpc"
	publish2 "DailyServer/back/logic_srv/publish"
	rpc2 "DailyServer/back/logic_srv/rpc"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/kafka"
	_ "github.com/micro/go-plugins/broker/kafka"
	"time"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example

// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE
func main() {
	microService := micro.NewService(micro.Name(constant.MICRO_LOGIC_SRV),
		micro.RegisterTTL(25*time.Second),
		micro.RegisterInterval(20*time.Second),
	)
	microService.Init(
		micro.Action(func(context *cli.Context) {
			config2.SetConf()
			glog2.InitLogger()
			db2.InitDb()
			if err := gredis2.InitRedis(); err != nil {
				glog2.Painc(err)
			}
		}),
	)

	micro.Broker(kafka.NewBroker(func(o *broker.Options) {
		o.Addrs = []string{
			"www.frozens.vip:9092",
		}
	}))

	broker.Init()

	if err := broker.Connect(); err != nil {
		glog2.Painc("Failed to connect kafka:", err)
		return
	}

	kafkaPub := new(publish2.KafkaPub)
	logicHandler := rpc2.NewLogicHandler(kafkaPub)
	err := grpc2.RegisterLogicServiceHandler(microService.Server(), logicHandler)
	if err != nil {
		glog2.Sugar().Panicf("RegisterLogicService is error %s", err.Error())
	}

	err = microService.Run()
	if err != nil {
		glog2.Sugar().Panicf("microServer run is error :%s", err)
	}

}
