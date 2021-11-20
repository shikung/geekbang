package main

import (
	"log"
	"net"
	pb "week4/api"
	"week4/internal/service"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, service.NewUserServcie())
	s.Serve(lis)
}
