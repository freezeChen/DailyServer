package main

import (
	"net"
	"net/rpc"
	"time"
)

var (
	logicRpcClient *rpc.Client

	loginRpcConnect = "RPC.Connect"
)

func InitLogicRPCClient(address string) (err error) {
	var conn net.Conn
	if conn, err = net.DialTimeout("tcp", address, 10*time.Second); err == nil {
		logicRpcClient = rpc.NewClient(conn)
	}
	return
}

func Connect(msg *Msg) (key string, err error) {
/*	var (
		arg   = &rpcEntity.ConnArg{Token: string("2")}
		reply = &rpcEntity.ConnReply{}
	)

	if logicRpcClient == nil {
		log.LogFile.E("logicRpcClient is nil")
	}

	if err = logicRpcClient.Call(loginRpcConnect, arg, reply); err != nil {
		log.LogFile.Error("Failed to call :%s error(%s)", loginRpcConnect, err)
		return
	}
	key = reply.Key*/
	return
}
