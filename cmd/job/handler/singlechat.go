/*
   @Time : 2019-01-31 11:10
   @Author : frozenchen
   @File : SingleChat
   @Software: DailyServer
*/
package handler

import (
	"DailyServer/grpc"
	"context"
)

type SingleChat struct {

}

func (SingleChat) Check(context.Context, *grpc.CheckReq, *grpc.CheckReply) error {
	panic("implement me")
}

