/*
   @Time : 2019-01-30 18:12
   @Author : frozenchen
   @File : main
   @Software: DailyServer
*/
package main

import (
	"DailyServer/commons/glog"
	"DailyServer/grpc"
	"DailyServer/lib"
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
)

var id = flag.Int64("id", 0, "help message.")
var toid = flag.Int64("tid", 0, "help message.")
var mm = flag.String("m", "", "help message.")

func main() {
	flag.Parse()
	glog.InitLogger()

	addr, _ := net.ResolveTCPAddr("tcp", "47.106.137.3:8020")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	rmsg := new(grpc.Proto)
	go func() {
		for {
			err := rmsg.ReadTCP(reader)
			if err != nil {
				panic(err)
			}
			var info lib.Info
			glog.Infof("proto:%+v", rmsg)
			json.Unmarshal(rmsg.Body, &info)

			glog.Infof("receive im info:%+v", info)

		}
	}()

	msg := new(grpc.Proto)
	msg.Ver = 1
	msg.Id = int32(*id)
	msg.Opr = grpc.OpAuth

	err = msg.WriteTCP(writer)
	if err != nil {
		panic(err)
	}

	msg.Opr = grpc.OpSendMsg

	info := new(lib.Info)
	info.Id = int32(*id)
	info.Rid = int32(*toid)
	info.Msg = *mm

	bytes, _ := json.Marshal(info)

	msg.Body = bytes

	err = msg.WriteTCP(writer)
	if err != nil {
		panic(err)
	}

	select {}
}
