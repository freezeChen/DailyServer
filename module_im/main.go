package main

import (
	"DailyServer/commons/glog"
	"net/http"
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
	http.HandleFunc("/ws", InitWebSocket)
	http.ListenAndServe(":8888", nil)

	InitSignal()
}
