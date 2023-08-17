package service

import (
	"context"
	"fmt"
	"github.com/2pgcn/auth/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/2pgcn/auth/api/auth/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type AuthService struct {
	pb.UnimplementedGreeterServer
	log  log.Logger
	user *biz.UserUsecase
}

func NewAuthService(log log.Logger, user *biz.UserUsecase) *AuthService {
	return &AuthService{
		log:  log,
		user: user,
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*emptypb.Empty, error) {
	//跳过验证
	return nil, s.user.Register(ctx, &biz.User{
		UserName: req.GetName(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
		Avatar:   req.GetAvatar(),
	})
}
func (s *AuthService) Login(ctx context.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	token, err := s.user.AuthByPwd(ctx, req.GetUsername(), req.GetPassword())
	fmt.Println(token, err)
	return &pb.AuthReply{Token: token}, err
}

func (s *AuthService) GetIdByToken(ctx context.Context, req *pb.GetIdByTokenReq) (*pb.GetIdByTokenReply, error) {
	id, err := s.user.GetIdByToken(ctx, req.GetToken())
	return &pb.GetIdByTokenReply{Id: id}, err
}

func (s *AuthService) UpdateUserInfo(ctx context.Context, req *pb.UpdateUserReq) (*emptypb.Empty, error) {
	return nil, nil
}
