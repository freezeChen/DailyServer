package im

import (
	"DailyServer/back/grpc"
	"bufio"
	"net"
)

//Channel
type Channel struct {
	Key    int32
	Ring   *Ring
	Reader bufio.Reader
	Writer bufio.Writer
	msg    chan *grpc.Proto
	Conn   *net.TCPConn
}

func NewChannel() (c *Channel) {
	c = new(Channel)
	c.msg = make(chan *grpc.Proto, 1024)
	c.Ring = InitRing()
	return c
}

func (c *Channel) Push(msg *grpc.Proto) (err error) {
	select {
	case c.msg <- msg:
	default:
	}
	return
}
func (c *Channel) Ready() *grpc.Proto {
	return <-c.msg
}

func (c *Channel) Signal() {
	c.msg <- grpc.ProtoReady
}

func (c *Channel) Close() {
	c.msg <- grpc.ProtoFinish
}
