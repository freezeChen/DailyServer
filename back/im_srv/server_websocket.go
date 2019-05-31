package im

import (
	"DailyServer/back/commons/glog"
	"DailyServer/back/grpc"
	lib2 "DailyServer/back/lib"
	"bufio"
	"context"
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

func InitWebSocket(srv *Server, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		glog.Error(err)
		return
	}
	ServeWebSocket(srv, conn)
}

func ServeWebSocket(srv *Server, ws *websocket.Conn) {
	var (
		err     error
		netConn = ws.UnderlyingConn()
		ch      = NewChannel()
		key     int32
		msg     *grpc.Proto
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	Reader = *bufio.NewReader(netConn)
	Writer = *bufio.NewWriter(netConn)

	if msg, err := Set(); err == nil {
		//连接合法性校验
		if key, err = srv.AuthWebSocket(ctx, ws, msg, ch); err == nil {
			Connect(ctx, key)
		}
	}

	if err != nil {
		Close()
		ws.Close()
		glog.Error(err)
	}

	go srv.dispatchWebsocket(ws, ch)

	for {
		if msg, err = Set(); err != nil {
			break
		}
		if err := msg.ReadWebSocket(ws); err != nil {
			break
		}
		SetAdv()
		Signal()
	}
}

func (srv *Server) dispatchWebsocket(ws *websocket.Conn, ch *Channel) {
	var err error
	for {
		msg := Ready()
		switch msg {
		case grpc.ProtoFinish:
			glog.Info("websocket finish")
			goto field1
		case grpc.ProtoReady:
			if msg, err = Get(); err != nil {
				goto field1
			}
			var info lib2.Info

			if err = json.Unmarshal(msg.Body, &info); err != nil {
				goto field1
			}

			if channel := Channel(msg.Toid); channel != nil {
				if err = Push(msg); err != nil {
					goto field1
				}
			}

			if err = msg.WriteWebSocket(ws); err != nil {
				goto field1
			}
			msg.Body = nil
			getAdv()
		default:

		}

	}

field1:
	glog.Errorf("dispatchTCP error(%s)", err)
	ws.Close()
	return

}

func (s *Server) AuthWebSocket(ctx context.Context, ws *websocket.Conn, msg *grpc.Proto, ch *Channel) (key int32, err error) {
	err = msg.ReadWebSocket(ws)
	if err != nil {
		return
	}
	key, err = AuthTCP(ctx, msg, ch)
	return
}
