/*
    @Time : 2018/8/31 下午2:28 
    @Author : 
    @File : microClient
    @Software: dailyserver2
*/
package main

import (
	"context"
	"dailyserver2/commons/glog"
	"dailyserver2/rpc/proto"
	"fmt"
	"github.com/micro/go-micro"
)

func InitMicroClient() {
	service := micro.NewService(micro.Name("module_user.client"))
	service.Init()

	userService := proto.NewUserService("module_user", service.Client())
	res, err := userService.Check(context.TODO(), &proto.CheckUserReq{Id: 1})
	if err != nil {
		glog.Debugf("failed to check:%s", err)
	}
fmt.Println(res.User)
	glog.Info(res.User)

}
