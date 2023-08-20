package data

import (
	"github.com/2pgcn/kratos-demo/apigw/internal/conf"
	v1 "github.com/2pgcn/kratos-demo/auth/api/auth/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	authClinet v1.
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	cleanup := func() {

		log.NewHelper(logger).Infof("closing the data resources,err:%v", err)
	}
	return &Data{rdb: redisClient, db: mysqlClient}, cleanup, nil
}
