package model

import (
	"bufio"
	"dailyserver/proto"
	"net"
)

type Channel struct {
	Id     int64
	Ring   *Ring
	Reader bufio.Reader
	Writer bufio.Writer
	msg    chan *proto.Proto
	Conn   *net.TCPConn
}

func NewChannel() (c *Channel) {
	c = new(Channel)
	c.msg = make(chan *proto.Proto, 1024)
	c.Ring = NewRing()
	return
}

func (c *Channel) Push(msg *proto.Proto) {
	select {
	case c.msg <- msg:
	default:

	}
	return
}

func (c *Channel) Ready() *proto.Proto {
	return <-c.msg
}

func (c *Channel) Signal() {
	c.msg <- proto.ProtoReady
}

func (c *Channel) Close() {
	c.msg <- proto.ProtoFinish
}
