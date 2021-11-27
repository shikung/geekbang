package main

import (
	"container/ring"
	"fmt"
)

var (
	head *ring.Ring // 环形队列（链表）
)

func main() {
	limitBucket := 12
	head = ring.New(limitBucket)
	for i := 0; i < limitBucket; i++ {
		head.Value = i
		head = head.Next()
	}
	fmt.Println("hello", head.Value)
	fmt.Println("hello", head.Len())
	fmt.Println("hello", head.Next().Value)
	fmt.Println("hello", head.Next().Next().Value)
	fmt.Println("hello", head.Next().Next().Next().Value)
	fmt.Println("hello", head.Next().Next().Next().Len())
}
