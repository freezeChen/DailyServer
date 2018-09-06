/*
    @Time : 2018/8/31 下午2:28 
    @Author : 
    @File : microClient
    @Software: DailyServer
*/
package main

import (
	"DailyServer/proto"
	"github.com/micro/go-micro"
)

var UserService proto.UserService

func InitMicroClient() {
	service := micro.NewService(micro.Name("module_user.client"))
	service.Init()
	UserService = proto.NewUserService("module_user", service.Client())
}
