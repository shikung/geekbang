package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func serverApp(ctx context.Context, stop <-chan string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "hello server app 5G")
	})
	server := http.Server{Addr: ":8090"}
	go func() {
		<-stop
		server.Shutdown(context.Background())
	}()
	server.ListenAndServe()
	return nil
}

func listenSignal(ctx context.Context, stop chan string) error {
	interrupt := make(chan os.Signal, 1)
	reload := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT)
	signal.Notify(reload, syscall.SIGHUP)
	for {
		select {
		case <-interrupt:
			fmt.Println("stop app")
			stop <- "stop"
			return errors.New("stop")
		case <-reload:
			fmt.Println("raload")
			return errors.New("reload")
		}
	}
}

func main() {
	fmt.Printf("welcom to app server \n")
	stop := make(chan string)
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return serverApp(ctx, stop)
	})
	group.Go(func() error {
		return listenSignal(ctx, stop)
	})
	err := group.Wait()
	fmt.Println(err)
	fmt.Printf("end app server begin\n")
	fmt.Println(ctx.Err())
}
