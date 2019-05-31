/*
   @Time : 2019-02-20 09:43
   @Author : frozenchen
   @File : kafkaPub
   @Software: DailyServer
*/
package publish

import (
	"DailyServer/back/constant"
	"DailyServer/back/grpc"
	"github.com/gogo/protobuf/proto"
	"github.com/micro/go-micro/broker"
)

type KafkaPub struct {
}

func (KafkaPub) PushSingleMsg(uid string, msg *grpc.Proto) error {

	bytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	err = broker.Publish(constant.JOB_TOPIC_SINGLECHAT, &broker.Message{
		Header: map[string]string{
			"UID": uid,
		},
		Body: bytes,
	})
	return err
}

func (KafkaPub) AuthSuccess(uid string) error {
	pro := new(grpc.Proto)
	pro.Ver = 1
	pro.Opr = grpc.OpAuthReply

	bytes, err := proto.Marshal(pro)
	if err != nil {
		return err
	}
	err = broker.Publish(constant.Job_Topic_AuthReply, &broker.Message{
		Header: map[string]string{
			"UID": uid,
		},
		Body: bytes,
	})

	return err
}
