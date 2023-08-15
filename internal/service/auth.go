package service

import (
	"context"
	"github.com/2pgcn/auth/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/2pgcn/auth/api/auth/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	pb.UnimplementedGreeterServer
	log  log.Logger
	User *biz.UserUsecase
}

func NewAuthService(log log.Logger, user *biz.UserUsecase) *AuthService {
	return &AuthService{
		log:  log,
		User: user,
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*emptypb.Empty, error) {
	return nil, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	return &pb.AuthReply{}, nil
}
func (s *AuthService) UpdateUserInfo(ctx context.Context, req *pb.UpdateUserReq) (*emptypb.Empty, error) {
	return nil, nil
}
