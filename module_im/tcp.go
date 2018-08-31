package main

import (
	"dailyserver2/commons/glog"
	"net"
	"fmt"
	"bufio"
)

func InitTCP() (err error) {
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

	BucketServer = NewBucket()

	go acceptTCP(Options, listener)
	return
}

func acceptTCP(options *options, listen *net.TCPListener) {
	var (
		conn *net.TCPConn
		err  error
	)
	for {
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

		go serverTCP(options, conn)
	}
}

func serverTCP(option *options, conn *net.TCPConn) {
	fmt.Println("tcp addr %s and %s", conn.LocalAddr(), conn.RemoteAddr())
	var (
		ch  *Channel
		msg *Msg
		key string
		err error
	)
	ch = NewChannel()
	ch.Reader = *bufio.NewReader(conn)
	ch.Writer = *bufio.NewWriter(conn)
	ch.Conn = conn

	if msg, err = ch.Ring.Set(); err == nil {
		//log.LogFile.I("msg :%s", msg)

		if key, err = AutoTCP(msg, ch); err == nil {
			BucketServer.Connect(key, ch)
		}
	}
	if err != nil {
		conn.Close()
		glog.Errorf("key:%s, handshake failed error")
	}

	go dispatchTCP(conn, ch)

	for {
		if msg, err = ch.Ring.Set(); err != nil {
			break
		}
		if err = msg.ReadTCP(&ch.Reader); err != nil {
			break
		}

		ch.Ring.SetAdv()
		ch.Signal()
	}
	conn.Close()
	ch.Close()

	fmt.Println("read is error(%s)", err)
}

func dispatchTCP(conn *net.TCPConn, ch *Channel) {
	var (
		err error
	)
	for {
		msg := ch.Ready()
		switch msg {
		case MsgFinish:
			goto field

		case MsgReady: //客户端发送过来的消息
			if msg, err = ch.Ring.Get(); err != nil {
				err = nil
				break
			}

			fmt.Println("msgready:", string(msg.Body))
			//log.LogFile.I("msgready: %s", msg)

			//todo  业务逻辑处理

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

func AutoTCP(msg *Msg, ch *Channel) (key string, err error) {
	key, err = Connect(msg)
	return
}

type TCPConfig struct {
	IP   string
	Port int
}
