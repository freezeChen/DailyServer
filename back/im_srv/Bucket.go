package im

import "sync"

type Bucket struct {
	lock sync.RWMutex
	chs  map[int32]*Channel
}

func NewBucket() (b *Bucket) {
	b = new(Bucket)
	b.chs = make(map[int32]*Channel)
	return b
}

func (b *Bucket) Online(key int32, c *Channel) {
	b.lock.Lock()
	//该id已登录 下线
	if ch, ok := b.chs[key]; ok {
		Close()
	}
	b.chs[key] = c
	//重复判断.
	b.lock.Unlock()
	return
}
func (b *Bucket) Get(key int32) (c *Channel) {
	b.lock.RLock()
	c = b.chs[key]
	b.lock.RUnlock()
	return
}

//下线&离线
func (b *Bucket) Offline(Key int32) {
	var (
		ok bool
	)
	b.lock.Lock()
	if _, ok = b.chs[Key]; ok {
		delete(b.chs, Key)
	}
	b.lock.Unlock()
}
