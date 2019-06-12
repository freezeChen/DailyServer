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
	xtime "dailyserver/lib/time"
	"dailyserver/proto"
	"math"
	"strconv"
	"time"

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

func (server *Server) serverTCP(c *conf.Config, conn *net.TCPConn, r int) {
	var (
		ch        *model.Channel
		msg       *proto.Proto
		err       error
		ctx       context.Context
		lastHb    = time.Now()
		timer     = server.Round.Timer(r)
		timerData *xtime.TimerData
	)

	ch = model.NewChannel()
	ch.Reader = *bufio.NewReader(conn)
	ch.Writer = *bufio.NewWriter(conn)
	ch.Conn = conn

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	timerData = timer.Add(time.Duration(_HandshakeTimeout*time.Second), func() {
		conn.Close()
	})

	if msg, err = ch.Ring.Set(); err == nil {
		if ch.Id, err = server.AuthTCP(ctx, msg, ch); err == nil {
			server.Bucket.Online(ch.Id, ch)
			zlog.Debugf("tcp connect id:%d proto: %+v", ch.Id, msg)
		}

		if err != nil {
			conn.Close()
			timer.Del(timerData)
			zlog.Errorf("id:%d handshake error(%v)", ch.Id, err)
			return
		}

		timerData.Key = strconv.FormatInt(ch.Id, 10)
		timer.Set(timerData, time.Duration(60*time.Second))
		go server.dispatchTCP(conn, &ch.Writer, ch)

		for {
			if msg, err = ch.Ring.Set(); err != nil {
				break
			}

			if err := msg.ReadTCP(&ch.Reader); err != nil {
				break
			}

			if msg.Opr == proto.OpHeartbeat {
				timer.Set(timerData, _HeartBeat)
				msg.Opr = proto.OpHeartbeatReply
				msg.Body = nil
				if now := time.Now(); now.Sub(lastHb) > _HeartBeat {
					if err = server.Heartbeat(ctx, ch.Id); err == nil {
						lastHb = now
					}
				}
			} else {
				if err := server.Operate(ctx, msg, ch); err != nil {
					break
				}
			}

			ch.Ring.SetAdv()
			ch.Signal()

		}
	} else {
		zlog.Errorf("set error(%s)", err)
	}

	timer.Del(timerData)
	conn.Close()
	ch.Close()

	if err := server.DisConnect(ctx, ch.Id); err != nil {
		zlog.Errorf("disConnect is error:(%v)", err)
	}
}

func (s *Server) dispatchTCP(conn *net.TCPConn, wr *bufio.Writer, ch *model.Channel) {
	var (
		err error
	)

	for {
		p := ch.Ready()
		switch p {
		case proto.ProtoFinish:
			goto failed
		case proto.ProtoReady:
			if p, err = ch.Ring.Get(); err != nil {
				break
			}

			if p.Opr == proto.OpHeartbeatReply {

			}

			if err := p.WriteTCP(wr); err != nil {
				goto failed
			}

			p.Body = nil
			ch.Ring.GetAdv()
		default:

		}

		if err := wr.Flush(); err != nil {
			goto failed
		}
	}

failed:
	conn.Close()
}
