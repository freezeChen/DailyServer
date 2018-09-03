/*
    @Time : 2018/9/3 下午2:10 
    @Author : 
    @File : imrpc
    @Software: dailyserver2
*/
package rpc

import (
	"dailyserver2/proto"
	"github.com/micro/go-micro"
)

var ImRPC proto.IMService

func InitImRPC() {
	service := micro.NewService(micro.Name("module_im.client"))
	service.Init()
	ImRPC = proto.NewIMService("module_im", service.Client())
}
