package main

import (
	"dailyserver2/commons/glog"
	"net"
	"net/http"
	"net/rpc"
)

func InitRPC(address string) (err error) {
	if err = rpc.Register(&PushRPC{}); err != nil {
		return
	}
	rpc.HandleHTTP()
	go RpcListen("tcp", address)
	return
}

func RpcListen(network string, addr string) {
	listener, err := net.Listen(network, addr)
	if err != nil {
		glog.Errorf("rpc listen failed: %s", err)
		panic(err)
	}
	defer func() {
		if err = listener.Close(); err != nil {
			glog.Errorf("Failed to close listener :%s", err)
		}
	}()

	http.Serve(listener, nil)

	//rpc.Accept(listener)
}

type PushRPC struct {
}

//func (rpc *PushRPC) PushMsg(arg *rpcEntity.PushMsgArg, reply *rpcEntity.MsgNoReply) error {
//	var (
//		ch  *Channel
//		err error
//	)
//	ch = BucketServer.Get(arg.Key)
//	err = ch.Push(arg.M)
//	return err
//}
