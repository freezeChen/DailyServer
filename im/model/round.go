/*
   @Time : 2019-06-05 14:15
   @Author : frozenchen
   @File : round
   @Software: DailyServer
*/
package model

import (
	"dailyserver/lib/time"
)

const (
	_timer     = 32
	_timerSize = 2048
)

// RoundOptions round options.
//type RoundOptions struct {
//	Timer        int
//	TimerSize    int
//	Reader       int
//	ReadBuf      int
//	ReadBufSize  int
//	Writer       int
//	WriteBuf     int
//	WriteBufSize int
//}

// Round userd for connection round-robin get a reader/writer/timer for split big lock.
type Round struct {
	timers []time.Timer
}

// NewRound new a round struct.
func NewRound() (r *Round) {
	var i int
	//r = &Round{
	//	options: RoundOptions{
	//		Reader:       c.TCP.Reader,
	//		ReadBuf:      c.TCP.ReadBuf,
	//		ReadBufSize:  c.TCP.ReadBufSize,
	//		Writer:       c.TCP.Writer,
	//		WriteBuf:     c.TCP.WriteBuf,
	//		WriteBufSize: c.TCP.WriteBufSize,
	//		Timer:        c.Protocol.Timer,
	//		TimerSize:    c.Protocol.TimerSize,
	//	}}
	// reader
	//r.readers = make([]bytes.Pool, r.options.Reader)
	//for i = 0; i < r.options.Reader; i++ {
	//	r.readers[i].Init(r.options.ReadBuf, r.options.ReadBufSize)
	//}
	//// writer
	//r.writers = make([]bytes.Pool, r.options.Writer)
	//for i = 0; i < r.options.Writer; i++ {
	//	r.writers[i].Init(r.options.WriteBuf, r.options.WriteBufSize)
	//}
	// timer

	r = &Round{}
	r.timers = make([]time.Timer, _timer)
	for i = 0; i < _timer; i++ {
		r.timers[i].Init(_timerSize)
	}
	return
}

// Timer get a timer.
func (r *Round) Timer(rn int) *time.Timer {
	return &(r.timers[rn%_timer])
}

//// Reader get a reader memory buffer.
//func (r *Round) Reader(rn int) *bytes.Pool {
//	return &(r.readers[rn%r.options.Reader])
//}
//
//// Writer get a writer memory buffer pool.
//func (r *Round) Writer(rn int) *bytes.Pool {
//	return &(r.writers[rn%r.options.Writer])
//}
