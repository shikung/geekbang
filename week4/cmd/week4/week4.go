package main

import (
	"log"
	"net"
	pb "week4/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) SayHello(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	return &pb.GetByUserID{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	s.Serve(lis)
}
