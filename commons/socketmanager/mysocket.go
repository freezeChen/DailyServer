package socketmanager

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"DailyServer/commons/log"
	"sync"
	"encoding/json"
)

type SocketManager struct {
	Lock sync.Mutex

	onlineSockets map[string]*onlineSocket
}

type onlineSocket struct {
	Conn *websocket.Conn
	Data chan interface{}
}

func NewOnlineSocket(conn *websocket.Conn, data chan interface{}) *onlineSocket {
	return &onlineSocket{Conn: conn, Data: data}

}

func HandelHttpToWebSocket(c *gin.Context) {
	server := &websocket.Server{Handler: handleWebSocket}
	server.ServeHTTP(c.Writer, c.Request)
}

func handleWebSocket(conn *websocket.Conn) {
	var b = make([]byte, 1024)
	var data string

	println(conn.Request())

	for {
		n, err := conn.Read(b)
		if err != nil {
			log.LogFile.I("failed to read data:%s", err)
		}

		data += string(b)

		if n == 0 {
			var temp interface{}

			json.Unmarshal([]byte(data), temp)
		}

	}

}
