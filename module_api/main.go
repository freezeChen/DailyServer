package main

import (
	"DailyServer/commons/config"
	"DailyServer/commons/db"
	"DailyServer/commons/glog"
	"DailyServer/commons/gredis"
	"DailyServer/module_api/routers"
	"DailyServer/module_api/rpc"
	"github.com/kardianos/service"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	SERVICE_NAME        string = "dailyserver"
	SERVICE_DISPLAYNAME string = "dailyserver"
	SERVICE_DESCRIPTION string = "本服务用于dailyserver的数据读取,请确保开机启动。"
)

type server struct{}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example

// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE
func main() {
	cfg := &service.Config{
		Name:        SERVICE_NAME,
		DisplayName: SERVICE_DISPLAYNAME,
		Description: SERVICE_DESCRIPTION,
	}
	s, err := service.New(&server{}, cfg)
	if err != nil {
		log.Fatalf("Failed to new s:%s\n", err)
	}

	if len(os.Args) == 2 {
		method := os.Args[1]
		err := service.Control(s, method)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if err = s.Run(); err != nil {
		log.Fatalf("Failed to run s:%s\n", err)
	}
}

func (ser server) Start(s service.Service) error {
	go ser.Run()
	return nil
}

func (s server) Run() {
	config.SetConf()
	glog.InitLogger()
	db.InitDb()
	gredis.InitRedis()
	rpc.InitRpc()

	r := routers.InitRouter()

	httpServer := http.Server{
		Addr:           ":8066",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	httpServer.ListenAndServe()
}

func (server) Stop(s service.Service) error {
	return nil
}
