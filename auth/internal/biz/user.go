package biz

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	v1 "github.com/2pgcn/kratos-demo/auth/api/auth/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type UserStatus int8

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUserInfo(ctx context.Context, userId int64, passwd string) error
	UpdateAvatar(ctx context.Context, userId int64, avatar string) error
	UpdateEmail(ctx context.Context, userId int64, email string) error
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByUserName(ctx context.Context, userName string) (*User, error)
	CreateTokenLinkId(ctx context.Context, id int64, token string) error
	GetIdByToken(ctx context.Context, token string) (int64, error)
}

type User struct {
	Id         int64      `json:"id"`
	UserName   string     `json:"user_name"`
	Password   string     `json:"password"`
	Email      string     `json:"email"`
	Avatar     string     `json:"avatar"`
	Status     UserStatus `json:"status"`
	CreateTime int64      `json:"create_time"`
	DeleteTime int64      `json:"delete_time"`
}

func (u *User) Encode() (string, error) {
	data, err := json.Marshal(u)
	return string(data), err
}

func (u *User) UnEncode(data string) error {
	return json.Unmarshal([]byte(data), &u)
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

// Register 创建并且生成token
func (u *UserUsecase) Register(ctx context.Context, user *User) error {
	return u.repo.CreateUser(ctx, user)
}

// 缓存错误暂不处理
func (u *UserUsecase) AuthByPwd(ctx context.Context, username string, passwd string) (string, error) {
	user, err := u.repo.FindByUserName(ctx, username)
	if err != nil {
		return "", v1.InternalError("FindIdByUserName err %v", err)
	}
	if user == nil {
		return "", errors.Unauthorized(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
	}
	if user.Password != passwd {
		return "", errors.Unauthorized(v1.ErrorReason_USER_AUTH_ERROR.String(), "password error")
	}
	token := createToken(user.Id)
	err = u.repo.CreateTokenLinkId(ctx, user.Id, token)
	return token, err
}

func (u *UserUsecase) GetIdByToken(ctx context.Context, token string) (int64, error) {
	id, err := u.repo.GetIdByToken(ctx, token)
	if id == 0 {
		return id, errors.Unauthorized(v1.ErrorReason_USER_AUTH_ERROR.String(), "token error")
	}
	return id, err
}

func createToken(id int64) string {
	data := md5.Sum([]byte(strconv.FormatInt(id, 10)))
	return hex.EncodeToString(data[:])
}
