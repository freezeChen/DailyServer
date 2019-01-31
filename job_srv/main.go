/*
   @Time : 2019-01-31 10:35
   @Author : frozenchen
   @File : main
   @Software: DailyServer
*/
package main

import (
	"DailyServer/constant"
	_ "github.com/micro/go-plugins/broker/kafka"
	"github.com/micro/go-micro"
)

func main() {
	microService := micro.NewService(
		micro.Name(constant.MICRO_JOB_SRV),
		micro.RegisterTTL(constant.MICRO_TTL),
		micro.RegisterInterval(constant.MICRO_Interval),
	)
	microService.Init()

	micro.Broker()

}
