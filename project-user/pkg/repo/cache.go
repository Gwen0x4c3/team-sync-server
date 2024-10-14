package repo

import (
	"context"
	"time"
)

// Cache 缓存接口
type Cache interface {
	Put(ctx context.Context, key string, value string, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
