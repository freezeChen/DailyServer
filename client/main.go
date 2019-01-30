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
	"fmt"
	"net"
)

func main() {

	glog.InitLogger()
	conn, err := net.Dial("tcp", "127.0.0.1:8020")
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

			json.Unmarshal(rmsg.Body, &info)

			fmt.Printf("info:%+v", info)

		}
	}()

	msg := new(grpc.Proto)
	msg.Opr = grpc.OpAuth
	msg.Body = []byte("5")

	err = msg.WriteTCP(writer)
	if err != nil {
		panic(err)
	}

	info := new(lib.Info)
	info.Id = 5
	info.Sid = 5
	info.Msg = "hello i am client"

	bytes, _ := json.Marshal(info)

	msg.Body = bytes

	msg.WriteTCP(writer)

}
