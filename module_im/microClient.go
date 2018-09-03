/*
    @Time : 2018/8/31 下午2:28 
    @Author : 
    @File : microClient
    @Software: dailyserver2
*/
package main

import (
	"dailyserver2/proto"
	"github.com/micro/go-micro"
)

var UserService proto.UserService

func InitMicroClient() {
	service := micro.NewService(micro.Name("module_user.client"))
	service.Init()
	UserService = proto.NewUserService("module_user", service.Client())
}
