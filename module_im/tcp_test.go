package main

import (
	"DailyServer/commons/glog"
	"DailyServer/lib"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"testing"
)

func TestS(t *testing.T) {

	glog.InitLogger()
	conn, err := net.Dial("tcp", "127.0.0.1:8020")
	if err != nil {
		t.Error(err)
	}

	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	rmsg := new(Msg)
	go func() {
		for {
			err := rmsg.ReadTCP(reader)
			if err != nil {
				t.Error(err)
			}
			var info lib.Info

			json.Unmarshal(rmsg.Body, &info)

			fmt.Printf("info:%+v", info)

		}
	}()

	msg := new(Msg)

	msg.Body = []byte("5")

	err = msg.WriteTCP(writer)
	if err != nil {
		t.Error(err)
	}

	info := new(lib.Info)
	info.Id = 5
	info.Sid = 5
	info.Msg = []byte("hello i am client")

	bytes, _ := json.Marshal(info)

	msg.Body = bytes

	msg.WriteTCP(writer)

	select {}
}

func TestA(t *testing.T) {
	var (
		num = 7
	)



	fmt.Println(num & 0)
	fmt.Println(num & 1)
	fmt.Println(num & 2)
	fmt.Println(num & 3)
	fmt.Println(num & 4)
	fmt.Println(num & 5)
	fmt.Println(num & 6)
	fmt.Println(num & 7)
	fmt.Println(num & 8)
	fmt.Println(num & 9)

}
