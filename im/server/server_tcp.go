/*
   @Time : 2019-05-31 15:23
   @Author : frozenchen
   @File : server_tcp
   @Software: DailyServer
*/
package server

import (
	"bufio"
	"context"
	"dailyserver/im/conf"

	"dailyserver/im/model"
	"dailyserver/proto"

	"github.com/freezeChen/studio-library/zlog"
	"net"
	"runtime"
)

func InitTCP(server *Server, c *conf.Config) (err error) {

	addr, err := net.ResolveTCPAddr("tcp", c.TCPPort)
	if err != nil {
		zlog.Errorf("net.ResolveTCPAddr error(%s)", err)
		return
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		zlog.Errorf("net.listenTCP error(%s)", err)
		return
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		go acceptTCP(server, c, listener)
	}

	return
}

func acceptTCP(server *Server, c *conf.Config, listen *net.TCPListener) {

	var (
		conn *net.TCPConn
		err  error
	)

	for {
		if conn, err = listen.AcceptTCP(); err != nil {
			zlog.Errorf("listen.AcceptTcp error(%s)", err)
			for {
				//等待连接
				if conn, err = listen.AcceptTCP(); err != nil {
					zlog.Errorf("listen.acceptTCP error(%s)", err)
					return
				}

				if err := conn.SetKeepAlive(c.TCPKeepalive); err != nil {
					zlog.Errorf("conn setKeepAlive error(%s)", err)
					return
				}

				if err := conn.SetReadBuffer(c.TCPReadBuffer); err != nil {
					zlog.Errorf("conn setReadBuffer error(%s)", err)
					return
				}

				if err := conn.SetWriteBuffer(c.TCPWriteBuffer); err != nil {
					zlog.Errorf("conn setWriterBuffer error(%s)", err)
					return
				}

				go server.serverTCP(c, conn)
			}
		}
	}
}

func (server *Server) serverTCP(c *conf.Config, conn *net.TCPConn) {
	var (
		ch  *model.Channel
		msg *proto.Proto
		err error
		ctx context.Context
	)

	ch = model.NewChannel()
	ch.Reader = *bufio.NewReader(conn)
	ch.Writer = *bufio.NewWriter(conn)
	ch.Conn = conn

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if msg, err = ch.Ring.Set(); err == nil {
		server.AuthTCP(ctx, msg, ch)
	}
	if err := conn.SetWriteBuffer(c.TCPWriteBuffer); err != nil {
		zlog.Errorf("conn setWriteBuffer error(%s)", err)
		return
	}

}
