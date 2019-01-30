package im

import (
	"DailyServer/commons/glog"
	"DailyServer/grpc"
	"DailyServer/lib"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"runtime"
)

func InitTCP(server *Server) (err error) {
	addr, err := net.ResolveTCPAddr("tcp", Options.TCPPort)
	if err != nil {
		glog.Error("net.ResolveTCPAddr error(%s)", err)
		return
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		glog.Error("tcp启动失败:", err)
		return
	}

	for i := 0; i < runtime.NumCPU(); i++ {
		go acceptTCP(server, Options, listener)
	}
	return
}

func acceptTCP(server *Server, options *options, listen *net.TCPListener) {
	var (
		conn *net.TCPConn
		err  error
	)
	for {
		//等待连接
		if conn, err = listen.AcceptTCP(); err != nil {
			glog.Error("listener acceptTCP error:", err)
			return
		}

		if err = conn.SetKeepAlive(options.TCPKeepalive); err != nil {
			glog.Error("listener setKeepAlive error:", err)
			return
		}

		if err = conn.SetReadBuffer(options.TCPReadBuffer); err != nil {
			glog.Error("listener set readBuffer error:", err)
			return
		}

		if err = conn.SetReadBuffer(options.TCPWriteBuffer); err != nil {
			glog.Errorf("listener set writeBuffer error:%s", err)
			return
		}

		go server.serverTCP(options, conn)
	}
}

func (s *Server) serverTCP(option *options, conn *net.TCPConn) {
	fmt.Printf("tcp addr %s and %s", conn.LocalAddr(), conn.RemoteAddr())

	var (
		ch    *Channel
		proto *grpc.Proto

		err error
		ctx context.Context
	)
	ch = NewChannel()
	ch.Reader = *bufio.NewReader(conn)
	ch.Writer = *bufio.NewWriter(conn)
	ch.Conn = conn
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if proto, err = ch.Ring.Set(); err == nil {
		//log.LogFile.I("proto :%s", proto)
		//连接合法性校验

		if ch.Key, err = s.AuthTCP(ctx, proto, ch); err == nil {
			s.Online(proto, ch)
		}
	}
	if err != nil {
		ch.Close()
		conn.Close()
		glog.Errorf("key:%s, handshake failed error", err)
		return
	}

	go s.dispatchTCP(conn, ch)

	for {
		if proto, err = ch.Ring.Set(); err != nil {
			break
		}
		if err = proto.ReadTCP(&ch.Reader); err != nil {
			break
		}

		ch.Ring.SetAdv()
		ch.Signal()
	}
	conn.Close()
	ch.Close()

	fmt.Printf("read is error(%s)", err)
}

func (s *Server) dispatchTCP(conn *net.TCPConn, ch *Channel) {
	var (
		err error
	)
	for {
		msg := ch.Ready()
		switch msg {
		case grpc.ProtoFinish:
			goto field

		case grpc.ProtoReady: //客户端发送过来的消息
			if msg, err = ch.Ring.Get(); err != nil {
				err = nil
				break
			}
			var info lib.Info
			if err = json.Unmarshal(msg.Body, &info); err != nil {
				goto field
			}

			if channel := s.Channel(info.Sid); channel != nil {
				err = channel.Push(msg)
				if err != nil {
					glog.Errorf("failed to push msg:%s", err)
				}
			}

			fmt.Println("msgready:", string(msg.Body))

			//将消息发送给自己
			if err = msg.WriteTCP(&ch.Writer); err != nil {
				goto field
			}
			msg.Body = nil
			ch.Ring.getAdv()

		default: //服务端返回给客户端的消息
			if err = msg.WriteTCP(&ch.Writer); err != nil {
				goto field
			}
		}
	}

field:
	glog.Errorf("dispatchTCP error(%s)", err)
	conn.Close()
	return
}

func (s *Server) AuthTCP(ctx context.Context, proto *grpc.Proto, ch *Channel) (key int32, err error) {
	err = proto.ReadTCP(&ch.Reader)
	if err != nil {
		glog.Error(err)
		return
	}

	if proto.Opr == grpc.OpAuth {
		key, err = s.Connect(ctx, proto.Id)
		if err != nil {
			glog.Errorf("tcp connect(id:%s).err:%v", proto.Id, key)
		}
	} else {
		err = grpc.ErrMsgNotCheck
	}

	return
}

type TCPConfig struct {
	IP   string
	Port int
}
