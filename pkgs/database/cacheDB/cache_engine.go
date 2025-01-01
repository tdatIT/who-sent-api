package cacheDB

import (
	"context"
	"github.com/tdatIT/who-sent-api/config"
	"time"
)

type CacheEngine interface {
	Get(ctx context.Context, key string) ([]byte, error)
	GetInt(ctx context.Context, key string) int
	Keys(ctx context.Context, pattern string) ([]string, error)
	Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	AddToSet(ctx context.Context, key string, val interface{}) error
	RemoveFromSet(ctx context.Context, key string, val interface{}) error
	ZAdd(ctx context.Context, key string, member interface{}, score float64) error
	ZPopMin(ctx context.Context, key string, count int) ([]string, error)
	ZRCard(ctx context.Context, key string) (int64, error)
	ZRange(ctx context.Context, key string, start, stop int64) ([]string, error)
	ZScore(ctx context.Context, keySet, member string) (float64, error)
	Expire(ctx context.Context, key string, ttl time.Duration) error
	Ping(ctx context.Context) error
	Close() error
}

func NewCacheEngine(cfg *config.AppConfig) (CacheEngine, error) {
	if cfg.Cache.Redis.Mode == "cluster" {
		return NewRedisCacheCluster(cfg)
	}
	return NewRedisCacheStandalone(cfg)
}
