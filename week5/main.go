package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"week5/athena"
)

func main() {
	ath := athena.NewAthena(time.Second*1, 10, 6)
	go ath.RefreshData()
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:9090") //获取一个tcpAddr
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr) //监听一个端口
	checkError(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept() // 在此处阻塞，每次来一个请求才往下运行handle函数
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(&conn, *ath) // 起一个单独的协程处理，有多少个请求，就起多少个协程，协程之间共享同一个全局变量limiting，对其进行原子操作。
	}
}

func handle(conn *net.Conn, ath athena.Athena) {
	defer (*conn).Close()
	ath.Add()
	//fmt.Println("handler n:", n)
	if ath.CurCount > int32(ath.LimitReq) { // 超出限频
		ath.Reset()
		(*conn).Write([]byte("HTTP/1.1 404 NOT FOUND\r\n\r\nError, too many request, please try again."))
	} else {
		ath.Cal()
		time.Sleep(1 * time.Second)                                             // 假设我们的应用处理业务用了1s的时间
		(*conn).Write([]byte("HTTP/1.1 200 OK\r\n\r\nI can change the world!")) // 业务处理结束后，回复200成功。
	}
}

// 异常报错的处理
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
