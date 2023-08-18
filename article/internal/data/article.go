package data

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"

	"github.com/2pgcn/kratos-demo/article/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/plugin/soft_delete"
)

type Article struct {
	Id         int64                 `gorm:"column:id;primaryKey;type:bigint;index;comment:id" json:"id"`
	Uid        int64                 `gorm:"column:uid;index;type:bigint;index;comment:user_id" json:"uid"`
	Title      string                `gorm:"column:title;type:varchar(32);comment:标题" json:"title"`
	Content    string                `gorm:"column:content;type:text;comment:博客内容" json:"content"`
	CreateTime int64                 `gorm:"column:create_time;type:bigint;not null" json:"create_time"`
	UpdateTime int64                 `gorm:"column:update_time;type:bigint;not null" json:"update_time"`
	DeleteTime soft_delete.DeletedAt `gorm:"column:delete_time;type:bigint;default:0;comment:删除时间" json:"-"`
}

func (a *Article) BeforeUpdate(tx *gorm.DB) error {
	if a.UpdateTime == 0 {
		a.UpdateTime = time.Now().Unix()
	}
	return nil
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	if a.CreateTime == 0 {
		a.CreateTime = time.Now().Unix()
	}
	if a.UpdateTime == 0 {
		a.UpdateTime = time.Now().Unix()
	}
	return nil
}

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *articleRepo) ListArticle(ctx context.Context, uid int64) ([]*biz.Article, error) {
	var lar []*Article
	err := ar.data.db.Model(&Article{}).Where("uid = ?", uid).Find(lar).Error
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Article, 0)
	for _, p := range lar {
		rv = append(rv, &biz.Article{
			ID:          p.Id,
			Uid:         p.Uid,
			Title:       p.Title,
			Content:     p.Content,
			CreatedTime: p.CreateTime,
			UpdatedTime: p.UpdateTime,
		})
	}
	return rv, nil
}

func (ar *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	var article *Article
	res := ar.data.db.Model(&Article{}).Where("id = ?", id).First(&article)
	if RecordNotFound(res) {
		return nil, nil
	}
	if res.Error != nil {
		return nil, res.Error
	}
	fmt.Println(article)
	return &biz.Article{
		ID:          article.Id,
		Title:       article.Title,
		Content:     article.Content,
		CreatedTime: article.CreateTime,
		UpdatedTime: article.UpdateTime,
	}, nil
}

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) (int64, error) {
	articleData := &Article{
		Uid:     article.Uid,
		Title:   article.Title,
		Content: article.Content,
	}
	rst := ar.data.db.WithContext(ctx).Model(&Article{}).Create(articleData)
	return articleData.Id, rst.Error
}

func (ar *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	err := ar.data.db.WithContext(ctx).Where("uid = ?", article.Uid).Model(&Article{}).Updates(&Article{
		Id:      id,
		Title:   article.Title,
		Content: article.Content,
	}).Error

	return err
}

func (ar *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	return ar.data.db.Model(&Article{}).Delete(&Article{}, id).Error
}
