package athena

import (
	"container/ring"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Athena struct {
	Size        time.Duration
	LimitReq    int
	limitBucket int
	CurCount    int32
	athena      *ring.Ring
}

func NewAthena(size time.Duration, limitReq int, limitBucket int) *Athena {
	athena := ring.New(limitBucket)
	for i := 0; i < limitBucket; i++ {
		athena.Value = 0
		athena = athena.Next()
	}
	ath := &Athena{
		Size:        size,
		LimitReq:    limitReq,
		limitBucket: limitBucket,
		CurCount:    0,
		athena:      athena,
	}
	return ath
}

// refresh data
func (a *Athena) RefreshData() {
	timer := time.NewTicker(a.Size)
	for range timer.C {
		subCount := int32(0 - a.athena.Value.(int))
		atomic.AddInt32(&a.CurCount, subCount)
		for i := 0; i < a.limitBucket; i++ {
			a.athena = a.athena.Next()
		}
		arr := [6]int{}
		for i := 0; i < a.limitBucket; i++ { // 这里是为了方便打印
			arr[i] = a.athena.Value.(int)
		}
		fmt.Println("move subCount,curCount,arr", subCount, a.CurCount, arr)
		a.athena.Value = 0
		a.athena = a.athena.Next()
	}
}

// Add a request
func (a *Athena) Add() int32 {
	return atomic.AddInt32(&a.CurCount, 1)
}

// Reset all request
func (a *Athena) Reset() {
	atomic.AddInt32(&a.CurCount, -1)
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
