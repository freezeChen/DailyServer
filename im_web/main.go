/*
    @Time : 2019-03-25 10:44 
    @Author : frozenchen
    @File : main
    @Software: DailyServer
*/
package main

import (
	"DailyServer/commons/config"
	"DailyServer/commons/db"
	"DailyServer/commons/glog"
	"DailyServer/commons/middleware"
	"DailyServer/constant"
	"DailyServer/im_web/controller"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
	"github.com/micro/go-web"
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
	controller.NewUserController().Router(router)
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
