package biz

import (
	"context"
	"time"
)

type Comment struct {
	ID           int32
	Content      string
	ContentId    int32
	ContentType  int8
	ContentTitle string
	Status       int8
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Like         int64
}

type CommentRepo interface {
	CreateComment(ctx context.Context, comment *Comment) (int32, error)
	GetComment(ctx context.Context, id int32) (*Comment, error)
}

type CommentUsercase struct {
	repo CommentRepo
}

func NewCommentUsercase(repo CommentRepo) *CommentUsercase {
	return &CommentUsercase{repo: repo}
}

func (uc *CommentUsercase) CreateComment(ctx context.Context, comment *Comment) (int32, error) {
	return uc.repo.CreateComment(ctx, comment)
}

func (uc *CommentUsercase) GetComment(ctx context.Context, id int32) (*Comment, error) {
	return uc.repo.GetComment(ctx, id)
}
