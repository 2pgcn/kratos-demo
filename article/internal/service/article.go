package service

import (
	"context"
	v1 "github.com/2pgcn/kratos-demo/article/api/article/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"go.opentelemetry.io/otel"

	pb "github.com/2pgcn/kratos-demo/article/api/article/v1"
	"github.com/2pgcn/kratos-demo/article/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

func NewBlogService(article *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{
		article: article,
		log:     log.NewHelper(logger),
	}
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	uid := ctx.Value("uid").(int64)
	aid, err := s.article.Create(ctx, &biz.Article{
		Uid:     uid,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return nil, errors.InternalServer(v1.ErrorReason_Internal_Error.String(), "内部基础错误")
	}
	return &pb.CreateArticleReply{
		Aid: aid,
	}, err
}

func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	err := s.article.Update(ctx, req.Id, &biz.Article{
		Uid:     ctx.Value("uid").(int64),
		Title:   req.Title,
		Content: req.Content,
	})
	return &pb.UpdateArticleReply{}, err
}

func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	err := s.article.Delete(ctx, req.Id)
	return &pb.DeleteArticleReply{}, err
}

func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	if req.Id < 1 {
		return nil, pb.ErrorBlogInvalidId("invalid article id")
	}
	tr := otel.Tracer("api")
	ctx, span := tr.Start(ctx, "GetArticle")
	defer span.End()
	p, err := s.article.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetArticleReply{Article: &pb.Article{Id: p.ID, Title: p.Title, Content: p.Content, CreateTime: p.CreatedTime}}, nil
}

func (s *ArticleService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	uid := ctx.Value("uid").(int64)
	ps, err := s.article.List(ctx, uid)
	reply := &pb.ListArticleReply{}
	for _, p := range ps {
		reply.Results = append(reply.Results, &pb.Article{
			Id:         p.ID,
			Title:      p.Title,
			Content:    p.Content,
			CreateTime: p.CreatedTime,
		})
	}
	return reply, err
}
