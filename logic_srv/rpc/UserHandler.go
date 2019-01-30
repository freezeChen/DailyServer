/*
    @Time : 2018/8/31 下午2:13 
    @Author : 
    @File : UserHandler
    @Software: DailyServer
*/
package rpc

import (
	"context"
	"DailyServer/commons/util"
	"DailyServer/logic_srv/models"
	"DailyServer/grpc"
	"fmt"
)

type UserHandler struct{}

func (UserHandler) Check(ctx context.Context, req *grpc.CheckUserReq, res *grpc.CheckUserRes) error {

	user, err := models.GetUserByID(util.ToInt64(req.Info))
	if err != nil {
		return err
	}

	//models.GetUserByID()
	fmt.Println(req.Info)
	res.User = &grpc.User{Id: user.Id}
	return nil
}
