package main

import (
	"bufio"
	"net"
)

//Channel
type Channel struct {
	id     string
	Ring   *Ring
	Reader bufio.Reader
	Writer bufio.Writer
	msg    chan *Msg
	Conn   *net.TCPConn
}

func NewChannel() (c *Channel) {
	c = new(Channel)
	c.msg = make(chan *Msg, 1024)
	c.Ring = InitRing()
	return c
}

func (c *Channel) Push(msg *Msg) (err error) {
	select {
	case c.msg <- msg:
	default:
	}
	return
}
func (c *Channel) Ready() *Msg {
	return <-c.msg
}

func (c *Channel) Signal() {
	c.msg <- MsgReady
}

func (c *Channel) Close() {
	c.msg <- MsgFinish
}
