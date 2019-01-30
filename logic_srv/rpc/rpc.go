/*
   @Time : 2018/8/31 下午2:22
   @Author : frozenChen
   @File : microServer
   @Software: DailyServer
*/
package rpc

import (
	"DailyServer/commons/glog"
	"DailyServer/constant"
	"DailyServer/grpc"
	"github.com/micro/go-micro"
	"time"
)

func InitRpc() {
	go func() {
		microS := micro.NewService(micro.Name(constant.MICRO_LOGIC_SRV),
			micro.RegisterTTL(25*time.Second),
			micro.RegisterInterval(20*time.Second),
		)
		microS.Init()

		grpc.RegisterLogicServiceHandler(microS.Server(), new(LogicHandler))
		err := microS.Run()
		if err != nil {
			glog.Sugar().Panicf("microServer run is error :%s", err)
		}
	}()
}
