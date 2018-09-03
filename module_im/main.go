package main

import (
	"dailyserver2/commons/glog"
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

	glog.Debugf("im RPC is running...")

	//glog.Debugf("logic server is running ...")

	InitSignal()
}
