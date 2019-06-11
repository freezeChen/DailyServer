/*
   @Time : 2019-01-30 18:12
   @Author : frozenchen
   @File : main
   @Software: DailyServer
*/
package main

import (
	"bufio"
	"dailyserver/proto"
	"fmt"
	"github.com/freezeChen/studio-library/zlog"
	"net"
	"time"
)

var scan string
//var tcpConn *net.TCPConn

func main() {

	zlog.InitLogger(&zlog.Config{
		Debug:      true,
		WriteKafka: false,
	})

	conn()

	return
main:
	fmt.Println(`请选择操作项目:
1:登录
2:单聊`)

	fmt.Scanln(&scan)

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

	conn, err := net.Dial("tcp", "127.0.0.1:8020")
	//addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8082")
	//tcpConn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)

	mWriter = bufio.NewWriter(conn)
	mReader = bufio.NewReader(conn)

	proto := new(proto.Proto)

	go func() {
		for {
			err := proto.ReadTCP(mReader)
			if err != nil {
				zlog.Error(err.Error())
				return
			}

			zlog.Infof("info :%+v", proto)

		}

	}()

	login()

}

func login() {

	msg := new(proto.Proto)
	msg.Ver = 1
	msg.Id = 1
	msg.Opr = proto.OpAuth

	err := msg.WriteTCP(mWriter)
	if err != nil {
		panic(err)
	}

	msg.Ver = 1
	msg.Id = 1
	msg.Toid = 2
	msg.Opr = proto.OpSendMsg
	msg.Body = []byte("hello i am client")

	msg.WriteTCP(mWriter)

	select {}

}
