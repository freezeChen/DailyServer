package main

import (
	"DailyServer/commons/config"
	"DailyServer/commons/db"
	"DailyServer/commons/glog"
	"DailyServer/commons/gredis"
	"DailyServer/constant"
	"DailyServer/grpc"
	"DailyServer/logic_srv/rpc"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
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
			gredis.InitRedis()
		}),
	)

	err := grpc.RegisterLogicServiceHandler(microService.Server(), new(rpc.LogicHandler))
	if err != nil {
		glog.Sugar().Panicf("RegisterLogicService is error %s", err.Error())
	}

	err = microService.Run()
	if err != nil {
		glog.Sugar().Panicf("microServer run is error :%s", err)
	}

}
