package main

import (
	"log"
	"net"
	"strings"

	pb "toupper/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Upper(ctx context.Context, in *pb.UpperRequest) (*pb.UpperReply, error) {
	log.Printf("Received: %s", in.Name)
	return &pb.UpperReply{Message: strings.ToUpper(in.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterToUpperServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
