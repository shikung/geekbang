package service

import (
	pb "week4/api"

	"golang.org/x/net/context"
)

func (s *UserService) GetUserById(ctx context.Context, req *pb.UserRequest) (*pb.UserReply, error) {

}
