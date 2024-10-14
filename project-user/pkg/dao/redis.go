package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Rc = &RedisCache{
		rdb: rdb,
	}
}

func (r *RedisCache) Put(ctx context.Context, key string, value string, expire time.Duration) error {
	err := r.rdb.Set(ctx, key, value, expire).Err()
	return err
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := r.rdb.Get(ctx, key).Result()
	return result, err
}
