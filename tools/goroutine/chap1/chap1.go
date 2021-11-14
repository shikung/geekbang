package chap1

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Add1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "hewwlo,GopherCOn 5G")
	})

	go func() {
		if err := http.ListenAndServe(":8090", nil); err != nil {
			log.Fatal(err)
		}
	}()
}

//没有地方给channel发送值，
//会导致接收操作一直是阻塞的，从而导致内存泄露
func Leak() {
	ch := make(chan int)
	go func() {
		val := <-ch
		fmt.Println("we received a value:", val)
	}()
}

func search(term string) (string, error) {
	time.Sleep(200 * time.Microsecond)
	return "same value", nil
}

func process(term string) error {
	record, err := search(term)
	if err != nil {
		return err
	}
	fmt.Println("Received:", record)
	return nil
}

func ServerApp() {
	fmt.Printf("hello Server app:" + "\n")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "hello,word")
	})
}

func ServerDebug() {
	http.ListenAndServe("127.0.0.1:8080", http.DefaultServeMux)
}
