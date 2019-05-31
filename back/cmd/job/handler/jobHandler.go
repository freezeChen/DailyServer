/*
   @Time : 2019-01-31 11:10
   @Author : frozenchen
   @File : SingleChat
   @Software: DailyServer
*/
package handler

import (
	glog2 "DailyServer/back/commons/glog"
	"DailyServer/back/constant"
	grpc2 "DailyServer/back/grpc"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/broker"
)

type JobHandler struct {
	imServer grpc2.IMService
}

func NewJobHandler(imServer grpc2.IMService) *JobHandler {
	return &JobHandler{imServer: imServer}
}

func (job JobHandler) Start() {

}

func (job JobHandler) SingChat() {
	_, err := broker.Subscribe(constant.JOB_TOPIC_SINGLECHAT, func(pub broker.Publication) error {
		glog2.Debug(constant.JOB_TOPIC_SINGLECHAT)
		var msg grpc2.Proto
		message := pub.Message()

		err := proto.Unmarshal(message.Body, &msg)
		if err != nil {
			glog2.Errorf("proto unmarshal:%s", err)
			return err
		}

		_, err = job.imServer.PushMsg(context.TODO(), &grpc2.PushMsgReq{Key: msg.Toid, Proto: &msg})
		if err != nil {
			glog2.Errorf("pushMsg:%s", err)
			return err
		}

		return nil
	})

	if err != nil {
		glog2.Painc(err)
	}
}

func (job JobHandler) AuthReply() {
	_, err := broker.Subscribe(constant.Job_Topic_AuthReply, func(pub broker.Publication) error {
		var pro grpc2.Proto
		message := pub.Message()
		fmt.Println("kafka receive authreply")
		err := proto.Unmarshal(message.Body, &pro)
		if err != nil {
			glog2.Errorf("proto unmarshal:%s", err)
			return err
		}

		_, err = job.imServer.PushMsg(context.TODO(), &grpc2.PushMsgReq{Key: pro.Id, Proto: &pro})

		return err
	})

	if err != nil {
		glog2.Painc(err)
	}
}
