package data

import (
	"context"
	"week4/internal/biz"
)

type UserRepo struct {
	data *Data
}

func NewUserRepo(data *Data) *UserRepo {
	return &UserRepo{data: data}
}

func (u *UserRepo) GetUser(ctx context.Context, uid int64) (*biz.User, error) {
	
}
