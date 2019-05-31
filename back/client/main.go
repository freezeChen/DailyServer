/*
   @Time : 2019-01-30 18:12
   @Author : frozenchen
   @File : main
   @Software: DailyServer
*/
package main

import (
	glog2 "DailyServer/back/commons/glog"
	util2 "DailyServer/back/commons/util"
	grpc2 "DailyServer/back/grpc"
	"bufio"
	"fmt"
	"net"
)

var scan string
var tcpConn *net.TCPConn

func main() {
	glog2.InitLogger()
main:
	fmt.Println(`请选择操作项目:
1:登录
2:单聊`)

	fmt.Scanln(&scan)
	fmt.Println(scan)

	switch scan {
	case "1":
		fmt.Println("请输入账号:")
		login()
	default:
		fmt.Println("选项错误")
		goto main
	}

	//reader := bufio.NewReader(os.Stdin)
	//data, _, _ := reader.ReadLine()
	//
	//
	//fmt.Println(string(data))

	//
	//flag.Parse()

	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//writer := bufio.NewWriter(tcpConn)
	//reader := bufio.NewReader(tcpConn)
	//rmsg := new(grpc.Proto)
	//go func() {
	//	for {
	//		err := rmsg.ReadTCP(reader)
	//		if err != nil {
	//			panic(err)
	//		}
	//		var info lib.Info
	//		glog.Infof("proto:%+v", rmsg)
	//		json.Unmarshal(rmsg.Body, &info)
	//
	//		glog.Infof("receive im info:%+v", info)
	//
	//	}
	//}()
	//
	//msg := new(grpc.Proto)
	//msg.Ver = 1
	//msg.Id = int32(*id)
	//msg.Opr = grpc.OpAuth
	//
	//err = msg.WriteTCP(writer)
	//if err != nil {
	//	panic(err)
	//}
	//
	//msg.Opr = grpc.OpSendMsg
	//
	//info := new(lib.Info)
	//info.Id = int32(*id)
	//info.Rid = int32(*toid)
	//info.Msg = *mm
	//
	//bytes, _ := json.Marshal(info)
	//
	//msg.Body = bytes
	//
	//err = msg.WriteTCP(writer)
	//if err != nil {
	//	panic(err)
	//}
	//
	select {}
}

var mWriter *bufio.Writer
var mReader *bufio.Reader

func conn() {
	var err error
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8020")
	tcpConn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	mWriter = bufio.NewWriter(tcpConn)
	mReader = bufio.NewReader(tcpConn)

	proto := new(grpc2.Proto)
	go func() {

		for {
			err := proto.ReadTCP(mReader)
			if err != nil {
				glog2.Error(err)
			}
			glog2.Infof("收到的信息:%+v", proto)

		}

	}()


}

func login() {
	fmt.Scanln(&scan)
	auth := int32(util2.ToInt(scan))

	conn()

	msg := new(grpc2.Proto)
	msg.Ver = 1
	msg.Id = auth
	msg.Opr = grpc2.OpAuth

	err := msg.WriteTCP(mWriter)
	if err != nil {
		panic(err)
	}

}
