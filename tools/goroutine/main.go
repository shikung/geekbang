package main

import (
	"fmt"
	"goroutine/chap1"
	"runtime"
	"sync"
	"time"
)

//func1
func func1() {
	go print("go Goroutine")
}

//func2
func func2() {
	name := "tian"
	go func() {
		fmt.Printf("hello,%s!\n", name)
	}()
	time.Sleep(time.Millisecond)
	name = "shikun"
}

//func3
func func3() {
	names := []string{"tianshikun", "wangming", "xiaoan", "ss"}
	for _, name := range names {
		go func(who string) {
			fmt.Printf("hello,%s!\n", who)
		}(name)
	}
	time.Sleep(time.Millisecond)
}

func func4() {
	fmt.Printf("runtime Compiler:%s!\n", runtime.Compiler)
	fmt.Printf("runtime GOARCH:%s!\n", runtime.GOARCH)
	fmt.Printf("runtime GOOS :%s!\n", runtime.GOOS)
	fmt.Printf("runtime GO ROOT:%s!\n", runtime.GOROOT())
}

func func5() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c <- "foo"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 60)
		fmt.Printf("message" + <-c)
	}()

	wg.Wait()
}

func func6() {
	c := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		c <- "foo"
		fmt.Printf("message" + "\n")
		time.Sleep(time.Second * 3)
		c <- "foo1"
		time.Sleep(time.Second * 3)
		fmt.Printf("message1" + "\n")
		c <- "foo2"
		time.Sleep(time.Second * 3)
		fmt.Printf("message2" + "\n")
		c <- "foo3"
		time.Sleep(time.Second * 3)
		fmt.Printf("message3" + "\n")
		c <- "foo4"
		time.Sleep(time.Second * 3)
		fmt.Printf("message4" + "\n")
		c <- "foo5"
		time.Sleep(time.Second * 3)
		fmt.Printf("message5" + "\n")
		c <- "foo6"
		time.Sleep(time.Second * 3)
		fmt.Printf("message6" + "\n")
		c <- "foo7"
		time.Sleep(time.Second * 3)
		fmt.Printf("message7" + "\n")
		c <- "foo8"
		time.Sleep(time.Second * 3)
		fmt.Printf("message8" + "\n")
		c <- "foo9"
		time.Sleep(time.Second * 3)
		fmt.Printf("message9" + "\n")
		c <- "foo10"
		time.Sleep(time.Second * 3)
		fmt.Printf("message10" + "\n")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 20)
		fmt.Printf("re message1" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message2" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message3" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message4" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message5" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message6" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message7" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message8" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message9" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message10" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message11" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message12" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message913" + <-c + "\n")
		time.Sleep(time.Second * 20)
		fmt.Printf("re message914" + <-c + "\n")
	}()

	wg.Wait()
}

func main() {
	chap1.Add1()
	chap1.ServerApp()
	func6()
}

type Resource struct {
	url        string
	polling    bool
	lastPolled int64
}

type Resources struct {
	data []*Resource
	lock *sync.Mutex
}

func Poller(res *Resources) {
	for {
		res.lock.Lock()
		var r *Resource
		for _, v := range res.data {
			if v.polling {
				continue
			}
			if r == nil || v.lastPolled < r.lastPolled {
				r = v
			}
		}
		if r != nil {
			r.polling = true
		}
		res.lock.Unlock()
		if r == nil {
			continue
		}
		res.lock.Lock()
		r.polling = false
		r.lastPolled = time.Microsecond.Nanoseconds()
		res.lock.Unlock()
	}
}

type Resource string

func Poller(in, out chan *Resource) {
	for r := range in {
		out <- r
	}
}
