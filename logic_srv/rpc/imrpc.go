/*
    @Time : 2018/9/3 下午2:10 
    @Author : 
    @File : imrpc
    @Software: DailyServer
*/
package rpc

import (
	"DailyServer/grpc"
	"github.com/micro/go-micro"
)

var ImRPC grpc.IMService

func InitImRPC() {
	service := micro.NewService(micro.Name("module_im.client"))
	service.Init()
	ImRPC = grpc.NewIMService("module_im", service.Client())
}
