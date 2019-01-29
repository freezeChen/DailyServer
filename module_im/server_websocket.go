package main

import (
	"DailyServer/commons/glog"
	"DailyServer/lib"
	"bufio"
	"encoding/json"
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

func ServeWebSocket(ws *websocket.Conn) {
	var (
		err     error
		netConn = ws.UnderlyingConn()
		ch      = NewChannel()
		key     int64
		msg     *Msg
	)
	ch.Reader = *bufio.NewReader(netConn)
	ch.Writer = *bufio.NewWriter(netConn)

	if msg, err := ch.Ring.Set(); err == nil {
		//连接合法性校验
		if key, err = AuthWebSocket(ws, msg); err == nil {
			BucketServer.Connect(key, ch)
		}
	}

	if err != nil {
		ch.Close()
		ws.Close()
		glog.Error(err)
	}

	go dispatchWebsocket(ws, ch)

	for {
		if msg, err = ch.Ring.Set(); err != nil {
			break
		}
		if err := msg.ReadWebSocket(ws); err != nil {
			break
		}
		ch.Ring.SetAdv()
		ch.Signal()
	}
}

func dispatchWebsocket(ws *websocket.Conn, ch *Channel) {
	var err error
	for {
		msg := ch.Ready()
		switch msg {
		case MsgFinish:
			glog.Info("websocket finish")
			goto field1
		case MsgReady:
			if msg, err = ch.Ring.Get(); err != nil {
				goto field1
			}
			var info lib.Info

			if err = json.Unmarshal(msg.Body, &info); err != nil {
				goto field1
			}

			if channel, ok := BucketServer.chs[info.Id]; ok {
				if err = channel.Push(msg); err != nil {
					goto field1
				}
			}

			if err = msg.WriteWebSocket(ws); err != nil {
				goto field1
			}
			msg.Body = nil
			ch.Ring.getAdv()
		}

	}

field1:
	glog.Errorf("dispatchTCP error(%s)", err)
	ws.Close()
	return

}

func AuthWebSocket(ws *websocket.Conn, msg *Msg) (key int64, err error) {
	err = msg.ReadWebSocket(ws)
	if err != nil {
		return
	}
	key, err = BucketServer.Operator.connect(msg)
	return
}
