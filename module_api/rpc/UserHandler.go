/*
    @Time : 2018/8/31 下午2:13 
    @Author : 
    @File : UserHandler
    @Software: dailyserver2
*/
package rpc

import (
	"context"
	"dailyserver2/commons/util"
	"dailyserver2/module_api/models"
	"dailyserver2/proto"
	"fmt"
)

type UserHandler struct{}

func (UserHandler) Check(ctx context.Context, req *proto.CheckUserReq, res *proto.CheckUserRes) error {

	user, err := models.GetUserByID(util.ToInt64(req.Info))
	if err != nil {
		return err
	}

	//models.GetUserByID()
	fmt.Println(req.Info)
	res.User = &proto.User{Id: user.Id}
	return nil
}
