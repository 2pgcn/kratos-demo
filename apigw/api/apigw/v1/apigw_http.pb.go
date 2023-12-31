// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v3.19.1
// source: apigw/v1/apigw.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationApigwCreateArticle = "/apigw.v1.Apigw/CreateArticle"
const OperationApigwDeleteArticle = "/apigw.v1.Apigw/DeleteArticle"
const OperationApigwGetArticle = "/apigw.v1.Apigw/GetArticle"
const OperationApigwListArticle = "/apigw.v1.Apigw/ListArticle"
const OperationApigwLogin = "/apigw.v1.Apigw/Login"
const OperationApigwRegister = "/apigw.v1.Apigw/Register"
const OperationApigwUpdateArticle = "/apigw.v1.Apigw/UpdateArticle"

type ApigwHTTPServer interface {
	CreateArticle(context.Context, *Article) (*CreateArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error)
	ListArticle(context.Context, *ListArticleRequest) (*ListArticleReply, error)
	Login(context.Context, *LoginRequest) (*ReturnTokenReply, error)
	Register(context.Context, *RegisterRequest) (*ReturnTokenReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleReply, error)
}

func RegisterApigwHTTPServer(s *http.Server, srv ApigwHTTPServer) {
	r := s.Route("/")
	r.POST("/pg/apigw/auth/v1", _Apigw_Register0_HTTP_Handler(srv))
	r.POST("/pg/apigw/auth/v1", _Apigw_Login0_HTTP_Handler(srv))
	r.POST("/pg/apigw/article/v1", _Apigw_CreateArticle0_HTTP_Handler(srv))
	r.GET("/pg/apigw/article/v1/{id}", _Apigw_GetArticle0_HTTP_Handler(srv))
	r.GET("/pg/apigw/article/v1", _Apigw_ListArticle0_HTTP_Handler(srv))
	r.PUT("/pg/apigw/article/v1/{id}", _Apigw_UpdateArticle0_HTTP_Handler(srv))
	r.DELETE("/pg/apigw/article/v1/{id}", _Apigw_DeleteArticle0_HTTP_Handler(srv))
}

func _Apigw_Register0_HTTP_Handler(srv ApigwHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApigwRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReturnTokenReply)
		return ctx.Result(200, reply)
	}
}

func _Apigw_Login0_HTTP_Handler(srv ApigwHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApigwLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ReturnTokenReply)
		return ctx.Result(200, reply)
	}
}

func _Apigw_CreateArticle0_HTTP_Handler(srv ApigwHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Article
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApigwCreateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateArticle(ctx, req.(*Article))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Apigw_GetArticle0_HTTP_Handler(srv ApigwHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApigwGetArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticle(ctx, req.(*GetArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Apigw_ListArticle0_HTTP_Handler(srv ApigwHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApigwListArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticle(ctx, req.(*ListArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Apigw_UpdateArticle0_HTTP_Handler(srv ApigwHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApigwUpdateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateArticle(ctx, req.(*UpdateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Apigw_DeleteArticle0_HTTP_Handler(srv ApigwHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApigwDeleteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*DeleteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteArticleReply)
		return ctx.Result(200, reply)
	}
}

type ApigwHTTPClient interface {
	CreateArticle(ctx context.Context, req *Article, opts ...http.CallOption) (rsp *CreateArticleReply, err error)
	DeleteArticle(ctx context.Context, req *DeleteArticleRequest, opts ...http.CallOption) (rsp *DeleteArticleReply, err error)
	GetArticle(ctx context.Context, req *GetArticleRequest, opts ...http.CallOption) (rsp *GetArticleReply, err error)
	ListArticle(ctx context.Context, req *ListArticleRequest, opts ...http.CallOption) (rsp *ListArticleReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *ReturnTokenReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *ReturnTokenReply, err error)
	UpdateArticle(ctx context.Context, req *UpdateArticleRequest, opts ...http.CallOption) (rsp *UpdateArticleReply, err error)
}

type ApigwHTTPClientImpl struct {
	cc *http.Client
}

func NewApigwHTTPClient(client *http.Client) ApigwHTTPClient {
	return &ApigwHTTPClientImpl{client}
}

func (c *ApigwHTTPClientImpl) CreateArticle(ctx context.Context, in *Article, opts ...http.CallOption) (*CreateArticleReply, error) {
	var out CreateArticleReply
	pattern := "/pg/apigw/article/v1"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApigwCreateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApigwHTTPClientImpl) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...http.CallOption) (*DeleteArticleReply, error) {
	var out DeleteArticleReply
	pattern := "/pg/apigw/article/v1/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApigwDeleteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApigwHTTPClientImpl) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...http.CallOption) (*GetArticleReply, error) {
	var out GetArticleReply
	pattern := "/pg/apigw/article/v1/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApigwGetArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApigwHTTPClientImpl) ListArticle(ctx context.Context, in *ListArticleRequest, opts ...http.CallOption) (*ListArticleReply, error) {
	var out ListArticleReply
	pattern := "/pg/apigw/article/v1"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationApigwListArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApigwHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*ReturnTokenReply, error) {
	var out ReturnTokenReply
	pattern := "/pg/apigw/auth/v1"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApigwLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApigwHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*ReturnTokenReply, error) {
	var out ReturnTokenReply
	pattern := "/pg/apigw/auth/v1"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApigwRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApigwHTTPClientImpl) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...http.CallOption) (*UpdateArticleReply, error) {
	var out UpdateArticleReply
	pattern := "/pg/apigw/article/v1/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApigwUpdateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
