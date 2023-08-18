package service

import (
	pb "github.com/2pgcn/kratos-demo/article/api/article/v1"
	"github.com/2pgcn/kratos-demo/article/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlogService)

type ArticleService struct {
	pb.UnimplementedArtictlServiceServer

	log *log.Helper

	article *biz.ArticleUsecase
}
