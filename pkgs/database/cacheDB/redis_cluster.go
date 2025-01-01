package cacheDB

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"github.com/tdatIT/who-sent-api/config"
	"github.com/tdatIT/who-sent-api/pkgs/logger"
	"time"
)

type redisCacheCluster struct {
	client *redis.ClusterClient
}

func NewRedisCacheCluster(cfg *config.AppConfig) (CacheEngine, error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: cfg.Cache.Redis.Address,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Fatalf("Failed to connect to redis: %v", err)
	}

	cacheEngine := redisCacheCluster{
		client: client,
	}
	return &cacheEngine, nil
}

func (r redisCacheCluster) Get(ctx context.Context, key string) ([]byte, error) {
	logger.Debugf("Get key: %s", key)
	result := r.client.Get(ctx, key)
	val, err := result.Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			logger.Debugf("Key not found: %s", key)
			return nil, nil
		}
		return nil, err
	}

	return val, err

}

func (r redisCacheCluster) GetInt(ctx context.Context, key string) int {
	logger.Debugf("Get key: %s", key)
	result := r.client.Get(ctx, key)
	val, err := result.Int()
	if err != nil {
		return 0
	}

	return val
}

func (r redisCacheCluster) Keys(ctx context.Context, pattern string) ([]string, error) {
	logger.Debugf("Get keys with pattern: %s", pattern)
	result := r.client.Keys(ctx, pattern)
	return result.Val(), result.Err()
}

func (r redisCacheCluster) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	logger.Debugf("Set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.Set(ctx, key, data, ttl)
	return result.Err()
}

func (r redisCacheCluster) Delete(ctx context.Context, key string) error {
	logger.Debugf("Deleting key: %s", key)
	result := r.client.Del(ctx, key)
	return result.Err()
}

func (r redisCacheCluster) AddToSet(ctx context.Context, key string, val interface{}) error {
	logger.Debugf("Add to set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.SAdd(ctx, key, data)
	return result.Err()
}

func (r redisCacheCluster) RemoveFromSet(ctx context.Context, key string, val interface{}) error {
	logger.Debugf("Remove from set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.SRem(ctx, key, data)
	return result.Err()
}

func (r redisCacheCluster) ZAdd(ctx context.Context, key string, member interface{}, score float64) error {
	logger.Debugf("ZAdd key: %s", key)
	data, err := sonic.Marshal(member)
	if err != nil {
		return err
	}

	result := r.client.ZAdd(ctx, key, redis.Z{Member: data, Score: score})
	return result.Err()
}

func (r redisCacheCluster) ZPopMin(ctx context.Context, key string, count int) ([]string, error) {
	logger.Debugf("ZPopMin key: %s", key)
	result := r.client.ZPopMin(ctx, key, int64(count))

	var values []string
	for _, z := range result.Val() {
		if member, ok := z.Member.(string); ok {
			values = append(values, member)
		}
	}

	return values, result.Err()
}

func (r redisCacheCluster) ZRCard(ctx context.Context, key string) (int64, error) {
	logger.Debugf("ZRCard key: %s", key)
	result := r.client.ZCard(ctx, key)
	return result.Val(), result.Err()
}

func (r redisCacheCluster) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	logger.Debugf("ZRange key: %s", key)
	result := r.client.ZRange(ctx, key, start, stop)
	return result.Val(), result.Err()
}

func (r redisCacheCluster) IsMember(ctx context.Context, key string, val interface{}) (bool, error) {
	logger.Debugf("Check member key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return false, err
	}

	result := r.client.SIsMember(ctx, key, data)
	return result.Val(), result.Err()
}

func (r redisCacheCluster) Expire(ctx context.Context, key string, ttl time.Duration) error {
	logger.Debugf("Expire key: %s", key)
	result := r.client.Expire(ctx, key, ttl)
	return result.Err()
}

func (r redisCacheCluster) ZScore(ctx context.Context, keySet, member string) (float64, error) {
	logger.Debugf("ZScore key: %s", keySet)
	result := r.client.ZScore(ctx, keySet, member)
	return result.Val(), result.Err()
}

func (r redisCacheCluster) Ping(ctx context.Context) error {
	logger.Debug("Ping redis")
	result := r.client.Ping(ctx)
	return result.Err()
}

func (r redisCacheCluster) Close() error {
	logger.Debugf("Closing redis connection host: %s", r.client.Options().Addrs)
	return r.client.Close()
}
