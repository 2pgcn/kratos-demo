package data

import (
	"context"
	"database/sql"
	"github.com/2pgcn/kratos-demo/article/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewArticleRepo)

// Data .
type Data struct {
	rdb *redis.Client
	db  *gorm.DB
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	redisClient, err := NewRedis(context.Background(), c.Redis)
	if err != nil {
		panic(err)
	}
	mysqlClient, err := NewMysql(c.Mysql)
	if err != nil {
		panic(err)
	}
	cleanup := func() {
		if err = redisClient.Close(); err != nil {
			log.NewHelper(logger).Errorf("closing the redis client err:%v", err)
		}
		var sqlDb *sql.DB
		if sqlDb, err = mysqlClient.DB(); err != nil {
			log.NewHelper(logger).Errorf("closing the redis client err:%v", err)
		}
		err = sqlDb.Close()
		log.NewHelper(logger).Infof("closing the data resources,err:%v", err)
	}
	return &Data{rdb: redisClient, db: mysqlClient}, cleanup, nil
}

func NewRedis(ctx context.Context, data *conf.Data_Redis) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:         data.Addr,
		Password:     "",              // no password set
		DB:           int(data.UseDb), // use default DB
		PoolSize:     int(data.PoolSize),
		ReadTimeout:  data.ReadTimeout.AsDuration(),
		WriteTimeout: data.WriteTimeout.AsDuration(),
	})
	client.Ping(ctx)
	return
}
