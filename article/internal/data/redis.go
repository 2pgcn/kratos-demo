package data

import (
	"context"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"strconv"
	"time"
)

var randTime = 60
var cacheMissVal = "-"

func (ar *articleRepo) GetArticleCache(ctx context.Context, id int64) (article string, err error) {
	get := ar.data.rdb.Get(ctx, articleKey(id))
	article, err = get.Result()
	if err == redis.Nil {
		return "", nil
	}
	return
}

func (ar *articleRepo) CreateArticleCache(ctx context.Context, id int64, data string) error {
	return ar.data.rdb.SetNX(ctx, articleKey(id), data, time.Hour+getRedisTmpNx()).Err()
}

func (ar *articleRepo) DeleteArticleCache(ctx context.Context, id int64) error {
	return ar.data.rdb.Del(ctx, articleKey(id)).Err()
}

func getRedisTmpNx() time.Duration {
	return time.Duration(rand.Intn(randTime)) * time.Second
}

func articleKey(id int64) string {
	return "pg:article:id:" + strconv.FormatInt(id, 10)
}
