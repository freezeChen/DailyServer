package socketmanager

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
)

func HandelHttpToWebSocket(c *gin.Context) {
	server := &websocket.Server{Handler: handleWebSocket}
	server.ServeHTTP(c.Writer, c.Request)
}

func handleWebSocket(conn *websocket.Conn) {

}
