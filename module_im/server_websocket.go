package main

import (
	"DailyServer/commons/glog"
	"bufio"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func InitWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Error(err)
		return
	}

	ServeWebSocket(conn)
}

func ServeWebSocket(conn *websocket.Conn) {
	var (
		err     error
		netConn = conn.UnderlyingConn()
		ch      = NewChannel()
		key     int64
		msg     *Msg
	)
	ch.Reader = *bufio.NewReader(netConn)
	ch.Writer = *bufio.NewWriter(netConn)

	if msg, err = ch.Ring.Set(); err == nil {
		//log.LogFile.I("msg :%s", msg)
		//连接合法性校验
		if key, err = AutoTCP(msg, ch); err == nil {
			BucketServer.Connect(key, ch)
		}
	}
	if err != nil {
		glog.Error(err)
	}

}
