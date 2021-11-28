package biz

import (
	"context"
)

type Comment struct {
	id   int32
	name string
}

type CommentRepo interface {
	CreateComment(context.Context, *Comment) error
	GetComment(context.Context, *Comment) error
}

type CommentUseCase struct {
	repo CommentRepo
}

func NewCommentUserCase(repo CommentRepo) *CommentUseCase {
	return &CommentUseCase{repo: repo}
}

func (uc *CommentUseCase) Create(ctx context.Context, m *Comment) error {
	return uc.repo.CreateComment(ctx, m)
}

func (uc *CommentUseCase) Get(ctx context.Context, m *Comment) error {
	return uc.repo.GetComment(ctx, m)
}
