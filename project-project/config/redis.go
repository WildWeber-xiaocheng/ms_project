package config

import (
	"github.com/WildWeber-xiaocheng/ms_project/project-project/internal/dao"
	"github.com/redis/go-redis/v9"
)

func (c *Config) ReConnRedis() {
	rdb := redis.NewClient(c.InitRedisOptions())
	rc := &dao.RedisCache{
		Rdb: rdb,
	}
	dao.Rc = rc
}
