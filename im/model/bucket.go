/*
   @Time : 2019-06-12 14:41
   @Author : frozenchen
   @File : bucket
   @Software: DailyServer
*/
package model

import "sync"

type Bucket struct {
	lock sync.RWMutex
	chs  map[int64]*Channel
}

func NewBucket() (b *Bucket) {
	b = new(Bucket)
	b.chs = make(map[int64]*Channel)
	return b
}

func (b *Bucket) Online(key int64, c *Channel) {
	b.lock.Lock()
	//该id已登录 下线
	if ch, ok := b.chs[key]; ok {
		ch.Close()
	}
	b.chs[key] = c
	//重复判断.
	b.lock.Unlock()
	return
}
func (b *Bucket) Get(key int64) (c *Channel) {
	b.lock.RLock()
	c = b.chs[key]
	b.lock.RUnlock()
	return
}

//下线&离线
func (b *Bucket) Offline(Key int64) {
	var (
		ok bool
	)
	b.lock.Lock()
	if _, ok = b.chs[Key]; ok {
		delete(b.chs, Key)
	}
	b.lock.Unlock()
}
