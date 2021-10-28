package main

import (
	"fmt"
	"runtime"
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
	fmt.Printf("go sys version:%s!\n", sys.TheVersion)
}

func main() {
	func4()
	func5()
	func3()
}
