/*
    @Time : 2018/8/31 下午2:13 
    @Author : 
    @File : UserHandler
    @Software: dailyserver2
*/
package rpchandle

import (
	"context"
	"dailyserver2/rpc/proto"
	"fmt"
)

type UserHandler struct{}

func (UserHandler) Check(ctx context.Context, req *proto.CheckUserReq, res *proto.CheckUserRes) error {
	fmt.Println(req)
	res.User = &proto.User{Id: 12346578}
	return nil
}
