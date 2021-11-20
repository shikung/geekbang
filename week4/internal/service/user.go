package service

import (
	pb "week4/api"

	"golang.org/x/net/context"
)

func (s *UserService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserReply, error) {
	reply := &pb.UserReply{}
	c, err := s.us.GetUser(ctx, req.Id)
	if err != nil {
		return reply, err
	}
	reply.Id = c.Id
	reply.Name = c.Name
	return reply, nil
}
