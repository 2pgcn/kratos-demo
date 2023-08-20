package biz

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
	ID          int64
	Uid         int64
	Title       string
	Content     string
	CreatedTime int64
	UpdatedTime int64
	Like        int64
}

var CacheMissVal = "-"

type ArticleRepo interface {
	// db
	ListArticle(ctx context.Context, uid int64) ([]*Article, error)
	GetArticle(ctx context.Context, id int64) (*Article, error)
	CreateArticle(ctx context.Context, article *Article) (int64, error)
	UpdateArticle(ctx context.Context, id int64, article *Article) error
	DeleteArticle(ctx context.Context, id int64) error

	// redis
	GetArticleCache(ctx context.Context, id int64) (article string, err error)
	CreateArticleCache(ctx context.Context, id int64, data string) error
	DeleteArticleCache(ctx context.Context, id int64) error
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

func (uc *ArticleUsecase) List(ctx context.Context, uid int64) (ps []*Article, err error) {
	ps, err = uc.repo.ListArticle(ctx, uid)
	if err != nil {
		return
	}
	return
}

func (uc *ArticleUsecase) Get(ctx context.Context, id int64) (p *Article, err error) {
	var arCache string
	if arCache, err = uc.repo.GetArticleCache(ctx, id); err != nil {
		log.Errorf("cache error %v", err)
	}
	if arCache == CacheMissVal {
		return nil, nil
	}
	if arCache != "" {
		err = json.Unmarshal([]byte(arCache), &p)
		return
	}
	p, err = uc.repo.GetArticle(ctx, id)
	if err != nil {
		return
	}
	if p == nil {
		_ = uc.repo.CreateArticleCache(ctx, id, CacheMissVal)
		return
	}
	var arCacheByte []byte
	arCacheByte, err = json.Marshal(&p)
	if err != nil {
		log.Errorf("cache error %v", err)
		return p, nil
	}
	err = uc.repo.CreateArticleCache(ctx, id, string(arCacheByte))
	if err != nil {
		log.Errorf("cache error %v", err)
		return p, nil
	}
	return
}

func (uc *ArticleUsecase) Create(ctx context.Context, article *Article) (int64, error) {
	return uc.repo.CreateArticle(ctx, article)
}

func (uc *ArticleUsecase) Update(ctx context.Context, id int64, article *Article) error {
	_ = uc.repo.DeleteArticleCache(ctx, id)
	return uc.repo.UpdateArticle(ctx, id, article)
}

func (uc *ArticleUsecase) Delete(ctx context.Context, id int64) error {
	_ = uc.repo.DeleteArticleCache(ctx, id)
	return uc.repo.DeleteArticle(ctx, id)
}
