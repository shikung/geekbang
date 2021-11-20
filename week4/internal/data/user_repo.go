package data

import (
	"context"
	"week4/internal/biz"
)

type User struct {
	id   int64
	name string
}

type UserRepo struct {
	data *Data
}

func NewUserRepo(data *Data) *UserRepo {
	return &UserRepo{data: data}
}

//
func (u *UserRepo) GetUser(ctx context.Context, uid int64) (*biz.User, error) {
	var user User
	err := u.data.QueryRow("SELECT id,name,title FROM test WHERE status=?", 1).Scan(&user.id, &user.name)
	if err != nil {
		return nil, err
	}
	return &biz.User{id: user.id, name: user.name}, nil
}
