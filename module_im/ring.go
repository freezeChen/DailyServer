package main

import "errors"

var (
	ErrRingEmpty = errors.New("ring buffer empty")
	ErrRingFull  = errors.New("ring buffer full")
)

//环形缓存区
type Ring struct {
	rp   uint64
	wp   uint64
	num  uint64
	data []Msg
}

func InitRing() (r *Ring) {
	r = new(Ring)
	r.num = 7
	r.data = make([]Msg, 8)
	return
}

func (r *Ring) Get() (msg *Msg, err error) {
	if r.rp == r.wp {
		return nil, ErrRingEmpty
	}
	msg = &r.data[r.rp&r.num]
	return
}

func (r *Ring) getAdv() {
	r.rp++
}

func (r *Ring) Set() (msg *Msg, err error) {
	if r.wp-r.rp >= r.num {
		return nil, ErrRingFull
	}
	msg = &r.data[r.wp&r.num]
	return
}

func (r *Ring) SetAdv() {
	r.wp++
}
