package main

import (
	"DailyServer/commons/config"
	"DailyServer/commons/db"
	"DailyServer/commons/glog"
	"DailyServer/commons/gredis"
	"DailyServer/constant"
	"DailyServer/grpc"
	"DailyServer/logic_srv/publish"
	"DailyServer/logic_srv/rpc"
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
			config.SetConf()
			glog.InitLogger()
			db.InitDb()
			if err := gredis.InitRedis(); err != nil {
				glog.Painc(err)
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
		glog.Painc("Failed to connect kafka:", err)
		return
	}

	kafkaPub := new(publish.KafkaPub)
	logicHandler := rpc.NewLogicHandler(kafkaPub)
	err := grpc.RegisterLogicServiceHandler(microService.Server(), logicHandler)
	if err != nil {
		glog.Sugar().Panicf("RegisterLogicService is error %s", err.Error())
	}

	err = microService.Run()
	if err != nil {
		glog.Sugar().Panicf("microServer run is error :%s", err)
	}

}
