/*
   @Time : 2019-01-31 11:10
   @Author : frozenchen
   @File : SingleChat
   @Software: DailyServer
*/
package handler

import (
	"DailyServer/commons/glog"
	"DailyServer/constant"
	"DailyServer/grpc"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/broker"
)

type JobHandler struct {
	imServer grpc.IMService
}

func NewJobHandler(imServer grpc.IMService) *JobHandler {
	return &JobHandler{imServer: imServer}
}

func (job JobHandler) Start() {

}

func (job JobHandler) SingChat() {
	_, err := broker.Subscribe(constant.JOB_TOPIC_SINGLECHAT, func(pub broker.Publication) error {
		glog.Debug(constant.JOB_TOPIC_SINGLECHAT)
		var msg grpc.Proto
		message := pub.Message()

		err := proto.Unmarshal(message.Body, &msg)
		if err != nil {
			glog.Errorf("proto unmarshal:%s", err)
			return err
		}

		_, err = job.imServer.PushMsg(context.TODO(), &grpc.PushMsgReq{Key: msg.Toid, Proto: &msg})
		if err != nil {
			glog.Errorf("pushMsg:%s", err)
			return err
		}

		return nil
	})

	if err != nil {
		glog.Painc(err)
	}
}

func (job JobHandler) AuthReply() {
	_, err := broker.Subscribe(constant.Job_Topic_AuthReply, func(pub broker.Publication) error {
		var pro grpc.Proto
		message := pub.Message()
		fmt.Println("kafka receive authreply")
		err := proto.Unmarshal(message.Body, &pro)
		if err != nil {
			glog.Errorf("proto unmarshal:%s", err)
			return err
		}

		_, err = job.imServer.PushMsg(context.TODO(), &grpc.PushMsgReq{Key: pro.Id, Proto: &pro})

		return err
	})

	if err != nil {
		glog.Painc(err)
	}
}
