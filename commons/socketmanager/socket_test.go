package socketmanager

import (
	"testing"
	"golang.org/x/net/websocket"
	"fmt"
	"DailyServer/api/models"
	"time"
	"encoding/json"
)

func TestName(t *testing.T) {

	ws, err := websocket.Dial("ws://127.0.0.1:8088/daily/socket", "", "http://127.0.0.1:8088/")
	if err != nil {
		fmt.Println(err)
	}

	time := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("time", time)

	info := models.Info{
		FUserID: "1",
		FInfo: models.ChatInfo{
			FTOID:    "2",
			FType:    1,
			FContent: "hello i am 1",
			FTime:    time,
		}}

	infostr, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(infostr))

	_, e := ws.Write(infostr)
	if e != nil {
		fmt.Println(e)
	}

	var temp string
	err = websocket.Message.Receive(ws, &temp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("server info :", temp)

}
