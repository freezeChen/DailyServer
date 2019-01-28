package main

import (
	"DailyServer/commons/glog"
	"DailyServer/lib"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"runtime"
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

	for i := 0; i < runtime.NumCPU(); i++ {
		go acceptTCP(Options, listener)
	}
	return
}

func acceptTCP(options *options, listen *net.TCPListener) {
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

		go serverTCP(options, conn)
	}
}

func serverTCP(option *options, conn *net.TCPConn) {
	fmt.Printf("tcp addr %s and %s", conn.LocalAddr(), conn.RemoteAddr())

	var (
		ch  *Channel
		msg *Msg
		key int64
		err error
	)
	ch = NewChannel()
	ch.Reader = *bufio.NewReader(conn)
	ch.Writer = *bufio.NewWriter(conn)
	ch.Conn = conn

	if msg, err = ch.Ring.Set(); err == nil {
		//log.LogFile.I("msg :%s", msg)
		//连接合法性校验
		if key, err = AutoTCP(msg, ch); err == nil {
			BucketServer.Connect(key, ch)
		}
	}
	if err != nil {
		ch.Close()
		conn.Close()
		glog.Errorf("key:%s, handshake failed error", err)
		return
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

	fmt.Printf("read is error(%s)", err)
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
			var info lib.Info
			if err = json.Unmarshal(msg.Body, &info); err != nil {
				goto field
			}

			if channel, ok := BucketServer.chs[info.Id]; ok {
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

func AutoTCP(msg *Msg, ch *Channel) (key int64, err error) {
	err = msg.ReadTCP(&ch.Reader)
	if err != nil {
		glog.Error(err)
		return
	}

	glog.Info("msg operation", msg.Operation)
	if msg.Operation == OPer_check {
		key, err = BucketServer.Operator.connect(msg)
	} else {
		err = ErrMsgNotCheck
	}

	return
}

type TCPConfig struct {
	IP   string
	Port int
}
