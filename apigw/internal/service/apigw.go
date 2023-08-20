package service

import (
	"context"

	pb "github.com/2pgcn/kratos-demo/apigw/api/apigw/v1"
)

type ApigwService struct {
	pb.UnimplementedApigwServer
}

func NewApigwService() *ApigwService {
	return &ApigwService{}
}

func (s *ApigwService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.ReturnTokenReply, error) {
	return &pb.ReturnTokenReply{}, nil
}
func (s *ApigwService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.ReturnTokenReply, error) {
	return &pb.ReturnTokenReply{}, nil
}
func (s *ApigwService) CreateArticle(ctx context.Context, req *pb.Article) (*pb.CreateArticleReply, error) {
	return &pb.CreateArticleReply{}, nil
}
func (s *ApigwService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	return &pb.GetArticleReply{}, nil
}
func (s *ApigwService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	return &pb.ListArticleReply{}, nil
}
func (s *ApigwService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	return &pb.UpdateArticleReply{}, nil
}
func (s *ApigwService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	return &pb.DeleteArticleReply{}, nil
}
