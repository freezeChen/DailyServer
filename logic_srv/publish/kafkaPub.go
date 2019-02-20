/*
   @Time : 2019-02-20 09:43
   @Author : frozenchen
   @File : kafkaPub
   @Software: DailyServer
*/
package publish

import (
	"DailyServer/constant"
	"DailyServer/grpc"
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
