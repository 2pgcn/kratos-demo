//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/2pgcn/kratos-demo/article/internal/biz"
	"github.com/2pgcn/kratos-demo/article/internal/conf"
	"github.com/2pgcn/kratos-demo/article/internal/data"
	"github.com/2pgcn/kratos-demo/article/internal/server"
	"github.com/2pgcn/kratos-demo/article/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
