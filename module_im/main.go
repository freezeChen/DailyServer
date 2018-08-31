package main

import (
	"dailyserver2/commons/glog"
)

func main() {
	glog.InitLogger()

	InitIMConfig()

	if err := InitTCP(); err != nil {
		panic(err)
	}

	/*if err := InitRPC(":8081"); err != nil {
		panic(err)
	}*/
	glog.Debugf("im RPC is running...")

	if err := InitLogicRPCClient("127.0.0.1" + ":8082"); err != nil {

		glog.Errorf("logincclient is error(%s)", err)
		//panic(err)
	}
	glog.Debugf("logic server is running ...")

	InitSignal()
}
