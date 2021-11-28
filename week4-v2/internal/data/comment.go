package data

import (
	"context"
	"week4-v2/internal/biz"
)

type CommentRepo struct {
	data *Data
}

func NewCommentRepo(data *Data) biz.CommentRepo {
	return &CommentRepo{
		data: data}
}

func (c *CommentRepo) CreateComment(ctx context.Context, g *biz.Comment) error {
	return nil
}

func (c *CommentRepo) GetComment(ctx context.Context, g *biz.Comment) error {
	return nil
}
