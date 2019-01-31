/*
   @Time : 2019-01-31 10:35
   @Author : frozenchen
   @File : main
   @Software: DailyServer
*/
package main

import (
	"DailyServer/commons/glog"
	"DailyServer/constant"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
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
				glog.InitLogger()

			},
		), )

	micro.Broker(kafka.NewBroker(func(options *broker.Options) {
		options.Addrs = []string{
			"test1:9092",
		}
	}))
	broker.Init()

	if err := broker.Connect(); err != nil {
		glog.Painc(err)
	}

	_, err := broker.Subscribe(constant.JOB_TOPIC_SINGLECHAT, func(pub broker.Publication) error {
		message := pub.Message()

		fmt.Print(string(message.Body))

		return nil
	})
	if err != nil {
		glog.Painc(err)
	}

	if err = microService.Run(); err != nil {
		glog.Painc(err)
	}

}
