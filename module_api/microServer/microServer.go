/*
    @Time : 2018/8/31 下午2:22 
    @Author : frozenChen
    @File : microServer
    @Software: dailyserver2
*/
package microServer

import (
	"dailyserver2/commons/glog"
	"dailyserver2/rpc/proto"
	"dailyserver2/rpc/rpchandle"
	"github.com/micro/go-micro"
)

func InitMicroServer() {
	go func() {
	
		microS := micro.NewService(micro.Name("module_user"))
		microS.Init()
		proto.RegisterUserServiceHandler(microS.Server(), new(rpchandle.UserHandler))
		err := microS.Run()
		if err != nil {
			glog.Sugar().Panicf("microServer run is error :%s", err)
		}
	}()
}
