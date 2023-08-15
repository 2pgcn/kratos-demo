package data

import (
	"context"
	"github.com/2pgcn/auth/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	redisClient *redis.Client
	mysqlClient *gorm.ConnPool
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	redisClient, err := NewRedis(context.Background(), c.Redis)
	if err != nil {
		panic(err)
	}

	cleanup := func() {
		if err = redisClient.Close(); err != nil {
			log.NewHelper(logger).Infof("closing the redis client err:%v", err)
		}
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{redisClient: redisClient}, cleanup, nil
}

func NewMysql() {

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

func NewKafka() {

}
