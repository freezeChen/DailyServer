package main

import (
	"DailyServer/commons/glog"
)

func main() {
	glog.InitLogger()
	InitIMConfig()
	if err := InitTCP(); err != nil {
		glog.Painc(err)
	}
	if err := InitRPC(); err != nil {
		glog.Painc(err)
	}
	InitMicroClient()
	glog.Debug("im RPC is running...")
	//glog.Debug("logic server is running ...")
	InitSignal()
}
