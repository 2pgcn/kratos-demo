package data

import (
	"context"

	"github.com/2pgcn/auth/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id         int            `gorm:"column:id;primaryKey;type:bigint;index;comment:数据id"`
	UserName   string         `gorm:"column:username;type:varchar(32);uniqueIndex;comment:用户名"`
	Password   string         `gorm:"column:password;type:varchar(32);comment:密码,纯md5,未加盐"`
	Email      string         `gorm:"column:email;type:varchar(64);index;comment:密码,纯md5,未加盐"`
	Avatar     string         `gorm:"column:avatar;type:text;comment:头像地址"`
	Status     biz.UserStatus `gorm:"column:status;type:tinyint;default:1"`
	CreateTime int32          `gorm:"column:create_time;type:int;not null"`
	DeleteTime int32          `gorm:"column:delete_time;type:int;default:0;comment:删除时间"`
}
type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewuserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	return nil
}
func (u *userRepo) UpdateUserInfo(ctx context.Context, userId int, passwd string) error {
	return nil
}
func (u *userRepo) UpdateAvatar(ctx context.Context, userId int, avatar string) error {
	return nil
}
func (u *userRepo) UpdateEmail(ctx context.Context, userId int, email string) error {
	return nil
}
func (u *userRepo) FindByID(ctx context.Context, id int64) (*biz.User, error) {
	return nil, nil
}
func (u *userRepo) FindByUserName(ctx context.Context, userName string) (*biz.User, error) {
	return nil, nil
}
