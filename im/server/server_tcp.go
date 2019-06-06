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
	"dailyserver/lib/time"
	"dailyserver/proto"
	"math"
	"strconv"
	time2 "time"

	"github.com/freezeChen/studio-library/zlog"
	"net"
	"runtime"
)

func (server *Server) InitTCP(c *conf.Config) (err error) {

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
		go server.acceptTCP(c, listener)
	}

	return
}

func (server *Server) acceptTCP(c *conf.Config, listen *net.TCPListener) {

	var (
		conn *net.TCPConn
		err  error
		r    int
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

				go server.serverTCP(c, conn, r)
				if r++; r == math.MaxInt32 {
					r = 0
				}
			}
		}
	}
}

func (server *Server) serverTCP(c *conf.Config, conn *net.TCPConn, r int) {
	var (
		ch  *model.Channel
		msg *proto.Proto
		err error
		ctx context.Context

		timer     = server.Round.Timer(r)
		timerData *time.TimerData
	)

	ch = model.NewChannel()
	ch.Reader = *bufio.NewReader(conn)
	ch.Writer = *bufio.NewWriter(conn)
	ch.Conn = conn

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timerData = timer.Add(time2.Duration(_HandshakeTimeout*time2.Second), func() {
		conn.Close()
	})

	if msg, err = ch.Ring.Set(); err == nil {
		if ch.Id, err = server.AuthTCP(ctx, msg, ch); err == nil {
			zlog.Debugf("tcp connect id:%d proto: %+v", ch.Id, msg)
		}

		if err != nil {
			conn.Close()
			timer.Del(timerData)
			zlog.Errorf("id:%d handshake error(%v)", err)
			return
		}

		timerData.Key = strconv.FormatInt(ch.Id, 10)
		timer.Set(timerData, time2.Duration(60*time2.Second))
		go server.dispatchTCP(conn, &ch.Writer, ch)
	}

}

func (s *Server) dispatchTCP(conn *net.TCPConn, wr *bufio.Writer, ch *model.Channel) {
	for {
		p := ch.Ready()
		switch p {
		case proto.ProtoFinish:
			goto failed
		case proto.ProtoReady:

		}
	}

failed:
}
