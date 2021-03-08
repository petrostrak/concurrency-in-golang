package main

import "sync"

// Barrier struct
type Barrier struct {
	total int
	count int
	mutex *sync.Mutex
	cond  *sync.Cond
}

// NewBarrier func
func NewBarrier(size int) *Barrier {
	lockToUse := &sync.Mutex{}
	condToUse := sync.NewCond(lockToUse)
	return &Barrier{size, size, lockToUse, condToUse}
}

// Wait func
func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count--
	if b.count == 0 {
		b.count = b.total
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.mutex.Unlock()
}
