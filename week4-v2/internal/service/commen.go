package service

import (
	"context"
	v1 "week4-v2/api/comment/v1"
	"week4-v2/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type CommentService struct {
	v1.UnimplementedCommentServer
	uc  *biz.CommentUseCase
	log *log.Helper
}

func NewCommentService(uc *biz.CommentUseCase) *CommentService {
	return &CommentService{uc: uc}
}

func (c *CommentService) CreateComment(ctx context.Context, req v1.CreateCommentRequest) (*v1.CreateCommentReply, error) {
	c.log.WithContext(ctx).Infof("ComementService")
	return &v1.CreateCommentReply{}, nil
}

func (c *CommentService) GetComment(ctx context.Context, req v1.GetCommentRequest) (*v1.GetCommentReply, error) {
	c.log.WithContext(ctx).Infof("ComementService")
	return &v1.GetCommentReply{}, nil
}
