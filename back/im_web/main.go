/*
    @Time : 2019-03-25 10:44 
    @Author : frozenchen
    @File : main
    @Software: DailyServer
*/
package main

import (
	"DailyServer/back/commons/config"
	"DailyServer/back/commons/db"
	"DailyServer/back/commons/glog"
	"DailyServer/back/commons/middleware"
	"DailyServer/back/constant"
	controller2 "DailyServer/back/im_web/controller"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
)

func main() {
	service := web.NewService(
		web.Address(":8021"),
		web.Name(constant.MICRO_IM_WEB),
		web.RegisterInterval(constant.MICRO_Interval),
		web.RegisterTTL(constant.MICRO_TTL))
	if err := service.Init(
		web.Action(func(context *cli.Context) {
			config.SetConf()
			glog.InitLogger()
			db.InitDb()
		})); err != nil {
		return
	}

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Static("/views", "./views")
	controller2.NewUserController().Router(router)
	if err := service.Init(); err != nil {
		panic(err)
		return
	}
	service.Handle("/", router)

	if err := service.Run(); err != nil {
		panic(err.Error())
		return
	}

}
