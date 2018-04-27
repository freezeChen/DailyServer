package socketmanager

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"sync"
	"fmt"
	"DailyServer/api/models"
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

	for {
		var temp models.Info
		var bytes []byte

		err := websocket.Message.Receive(conn, &bytes)
		if err != nil {
			fmt.Println(err)
		}

		err = json.Unmarshal(bytes, &temp)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("message:", temp.FUserID, " : ", temp.FInfo.FContent)
	}

}
