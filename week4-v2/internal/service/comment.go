package service

import (
	"context"
	"fmt"

	pb "week4-v2/api/comment/v1"
	"week4-v2/internal/biz"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
	comment *biz.CommentUsercase
}

func NewCommentService(comment *biz.CommentUsercase) *CommentService {
	return &CommentService{
		comment: comment,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	fmt.Print(req.ContentId)
	fmt.Print(req.ContentTitle)
	fmt.Print(req.ContentType)
	id, err := s.comment.CreateComment(ctx, &biz.Comment{
		Content:      req.Content,
		ContentId:    req.ContentId,
		ContentTitle: req.ContentTitle,
		ContentType:  int8(req.ContentType),
	})
	return &pb.CreateCommentReply{Id: id}, err
}

func (s *CommentService) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentReply, error) {
	res, err := s.comment.GetComment(ctx, req.Id)
	return &pb.GetCommentReply{
		Comment: &pb.Comment{Id: res.ID, Content: res.Content, ContentId: res.ContentId, ContentType: int32(res.ContentType), ContentTitle: res.ContentTitle}}, err
}
