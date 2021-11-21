package athena

import (
	"container/ring"
	"sync"
	"sync/atomic"
	"time"
)

type Athena struct {
	limitReq    int
	limitBucket int
	curCount    int32
	athena      *ring.Ring
}

// refresh data
func (a *Athena) RefreshData() {
	timer := time.NewTicker(time.Microsecond * 100)
	for range timer.C {
		subCount := int32(0 - a.athena.Value.(int))
		atomic.AddInt32(&a.curCount, subCount)
		for i := 0; i < a.limitBucket; i++ {
			a.athena = a.athena.Next()
		}
		a.athena.Value = 0
		a.athena = a.athena.Next()
	}
}

// Add a request
func (a *Athena) Add() int32 {
	return atomic.AddInt32(&a.curCount, 1)
}

// Reset all request
func (a *Athena) Reset() {
	atomic.AddInt32(&a.curCount, -1)
}

// Cal request
func (a *Athena) Cal() {
	mu := sync.Mutex{}
	mu.Lock()
	pos := a.athena.Prev()
	val := pos.Value.(int)
	val++
	pos.Value = val
	mu.Unlock()
}
