package config

import (
	"github.com/redis/go-redis/v9"
	"test.com/project-project/internal/dao"
)

func (c *Config) ReConnRedis() {
	rdb := redis.NewClient(c.InitRedisOptions())
	rc := &dao.RedisCache{
		Rdb: rdb,
	}
	dao.Rc = rc
}
