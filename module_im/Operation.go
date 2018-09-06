package main

import (
	"context"
	"DailyServer/proto"
)

type Operator interface {
	connect(msg *Msg) (key int64, err error)
}

type DefaultOperator struct {
}

func (o *DefaultOperator) connect(msg *Msg) (key int64, err error) {
	bytes, _ := msg.Body.MarshalJSON()
	userRes, err := UserService.Check(context.TODO(), &proto.CheckUserReq{Info: string(bytes)})
	if err == nil {
		key = userRes.User.Id
	}

	return
}
