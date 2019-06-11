/*
   @Time : 2019-06-10 13:42:13
   @Author :
   @File : service
   @Software: logic
*/
package service

import (
	"context"
	"dailyserver/proto"
)

func (Service) Auth(ctx context.Context, req *proto.AuthReq, reply *proto.EmptyReply) error {
	return nil
}
