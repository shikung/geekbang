package service

import (
	pb "week4/api"
	"week4/internal/biz"
)

type UserService struct {
	pb.UnimplementedCartServer

	us *biz.UserCase
}

func NewUserServcie(us *biz.UserCase) *UserService {
	return &UserService{us: us}
}
