package dao

import (
	"github.com/Gwen0x4c3/team-sync-server/project-user/config"
	_ "github.com/Gwen0x4c3/team-sync-server/project-user/config"

	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Addr,
		Password: config.Cfg.Redis.Password,
		DB:       config.Cfg.Redis.DB,
	})
	Rc = &RedisCache{
		rdb: rdb,
	}
	log.Println("Init redis cache")
}

func (r *RedisCache) Put(ctx context.Context, key string, value string, expire time.Duration) error {
	err := r.rdb.Set(ctx, key, value, expire).Err()
	return err
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := r.rdb.Get(ctx, key).Result()
	return result, err
}
