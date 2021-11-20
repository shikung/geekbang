package biz

import "context"

type User struct {
	Id   int64
	Name string
}

type UserCase struct {
	repo CartRepo
}

type CartRepo interface {
	GetUser(ctx context.Context, uid int64) (*User, error)
}

func NewUserCase(repo CartRepo) *UserCase {
	return &UserCase{repo: repo}
}

func (us *UserCase) GetUser(ctx context.Context, uid int64) (*User, error) {
	return us.repo.GetUser(ctx, uid)
}
