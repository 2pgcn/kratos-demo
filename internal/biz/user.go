package biz

import (
	"context"
	v1 "github.com/2pgcn/auth/api/auth/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type UserStatus int8

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUserInfo(ctx context.Context, userId int, passwd string) error
	UpdateAvatar(ctx context.Context, userId int, avatar string) error
	UpdateEmail(ctx context.Context, userId int, email string) error
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByUserName(ctx context.Context, userName string) (*User, error)
}

type User struct {
	Id         int
	UserName   string
	Password   string
	Email      string
	Avatar     string
	Status     UserStatus
	CreateTime int32
	DeleteTime int32
}

// GreeterUsecase is a Greeter usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewUsercase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}
