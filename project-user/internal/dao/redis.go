package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"test.com/project-user/config"
	"time"
)

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() { //redis连接
	rdb := redis.NewClient(config.Conf.InitRedisOptions())
	Rc = &RedisCache{
		rdb: rdb,
	}
}

func (rc *RedisCache) Put(ctx context.Context, key string, value string, expire time.Duration) error {
	err := rc.rdb.Set(ctx, key, value, expire).Err()
	return err
}

func (rc *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := rc.rdb.Get(ctx, key).Result()
	return result, err
}
