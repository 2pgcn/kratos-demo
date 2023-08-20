package data

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"

	"github.com/2pgcn/kratos-demo/auth/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Id         int64          `gorm:"column:id;primaryKey;type:bigint;index;comment:数据id"`
	UserName   string         `gorm:"column:username;type:varchar(32);uniqueIndex;comment:用户名"`
	Password   string         `gorm:"column:password;type:varchar(32);comment:密码,纯md5,未加盐"`
	Email      string         `gorm:"column:email;type:varchar(64);index;comment:密码,纯md5,未加盐"`
	Avatar     string         `gorm:"column:avatar;type:text;comment:头像地址"`
	Status     biz.UserStatus `gorm:"column:status;type:tinyint;default:1"`
	CreateTime int64          `gorm:"column:create_time;type:bigint;not null"`
	UpdateTime int64          `gorm:"column:update_time;type:bigint;not null"`
	DeleteTime int64          `gorm:"column:delete_time;type:bigint;default:0;comment:删除时间"`
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if u.UpdateTime == 0 {
		u.UpdateTime = time.Now().Unix()
	}
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.CreateTime == 0 {
		u.CreateTime = time.Now().Unix()
	}
	if u.UpdateTime == 0 {
		u.UpdateTime = time.Now().Unix()
	}
	return nil
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
	userPo := &User{
		UserName:   user.UserName,
		Password:   user.Password,
		Email:      user.Email,
		Avatar:     user.Avatar,
		DeleteTime: 0,
	}
	return u.data.mysqlClient.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
	}).Save(&userPo).Error
}

func (u *userRepo) UpdateUserInfo(ctx context.Context, userId int64, passwd string) error {
	return u.data.mysqlClient.WithContext(ctx).Model(&User{}).Where("id = ?", userId).Update("password", passwd).Error
}

func (u *userRepo) UpdateAvatar(ctx context.Context, userId int64, avatar string) error {
	return u.data.mysqlClient.WithContext(ctx).Model(&User{}).Where("id = ?", userId).Update("avatar", avatar).Error
}

func (u *userRepo) UpdateEmail(ctx context.Context, userId int64, email string) error {
	return u.data.mysqlClient.WithContext(ctx).Model(&User{}).Where("id = ?", userId).Update("email", email).Error
}

func (u *userRepo) FindByID(ctx context.Context, id int64) (*biz.User, error) {
	var user *User
	err := u.data.mysqlClient.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:         user.Id,
		UserName:   user.UserName,
		Password:   user.Password,
		Email:      user.Email,
		Avatar:     user.Avatar,
		Status:     user.Status,
		CreateTime: user.CreateTime,
		DeleteTime: user.DeleteTime,
	}, nil
}

func (u *userRepo) FindByIDCache(ctx context.Context, id int64) (user *biz.User, err error) {
	rst := u.data.redisClient.Get(ctx, string(id))
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = user.UnEncode(rst.String())
	return

}

func (u *userRepo) FindByUserName(ctx context.Context, userName string) (*biz.User, error) {
	var user *User
	err := u.data.mysqlClient.WithContext(ctx).Where("username = ?", userName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:         user.Id,
		UserName:   user.UserName,
		Password:   user.Password,
		Email:      user.Email,
		Avatar:     user.Avatar,
		Status:     user.Status,
		CreateTime: user.CreateTime,
		DeleteTime: user.DeleteTime,
	}, nil
}

func (u *userRepo) CreateTokenLinkId(ctx context.Context, id int64, token string) error {
	return u.data.redisClient.SetNX(ctx, getRedisUserTokenKey(token), id, time.Hour*3).Err()
}

func (u *userRepo) GetIdByToken(ctx context.Context, token string) (int64, error) {
	fmt.Println(getRedisUserTokenKey(token))
	rst := u.data.redisClient.Get(ctx, getRedisUserTokenKey(token))
	id, err := rst.Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return id, err
}

func getRedisUserTokenKey(token string) string {
	return "pg:auth:token:" + token
}
