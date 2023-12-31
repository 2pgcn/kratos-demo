// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/2pgcn/kratos-demo/auth/internal/biz"
	"github.com/2pgcn/kratos-demo/auth/internal/conf"
	"github.com/2pgcn/kratos-demo/auth/internal/data"
	"github.com/2pgcn/kratos-demo/auth/internal/server"
	"github.com/2pgcn/kratos-demo/auth/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUsercase(userRepo, logger)
	authService := service.NewAuthService(logger, userUsecase)
	grpcServer := server.NewGRPCServer(confServer, authService, logger)
	httpServer := server.NewHTTPServer(confServer, authService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
