package cacheDB

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"github.com/tdatIT/who-sent-api/config"
	"github.com/tdatIT/who-sent-api/pkgs/logger"

	"strings"

	"time"
)

type redisCacheStandalone struct {
	client *redis.Client
}

func NewRedisCacheStandalone(cfg *config.AppConfig) (CacheEngine, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Cache.Redis.Address[0],
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		logger.Fatalf("Failed to connect to redis: %v", err)
	}

	cacheEngine := redisCacheStandalone{
		client: client,
	}
	return &cacheEngine, nil
}

func (r redisCacheStandalone) Get(ctx context.Context, key string) ([]byte, error) {
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

func (r redisCacheStandalone) GetInt(ctx context.Context, key string) int {
	logger.Debugf("Get key: %s", key)
	result := r.client.Get(ctx, key)
	val, err := result.Int()
	if err != nil {
		return 0
	}

	return val
}
func (r redisCacheStandalone) Keys(ctx context.Context, pattern string) ([]string, error) {
	logger.Debugf("Get keys with pattern: %s", pattern)
	result := r.client.Keys(ctx, pattern)
	return result.Val(), result.Err()
}

func (r redisCacheStandalone) Set(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	logger.Debugf("Set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.Set(ctx, key, data, ttl)
	return result.Err()
}

func (r redisCacheStandalone) AddToSet(ctx context.Context, key string, val interface{}) error {
	logger.Debugf("Add to set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.SAdd(ctx, key, data)
	return result.Err()
}

func (r redisCacheStandalone) RemoveFromSet(ctx context.Context, key string, val interface{}) error {
	logger.Debugf("Remove from set key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return err
	}

	result := r.client.SRem(ctx, key, data)
	return result.Err()
}

func (r redisCacheStandalone) IsMember(ctx context.Context, key string, val interface{}) (bool, error) {
	logger.Debugf("Check member key: %s", key)
	data, err := sonic.Marshal(val)
	if err != nil {
		return false, err
	}

	result := r.client.SIsMember(ctx, key, data)
	return result.Val(), result.Err()
}

func (r redisCacheStandalone) Delete(ctx context.Context, key string) error {
	logger.Debugf("Deleting key: %s", key)
	result := r.client.Del(ctx, key)
	return result.Err()
}

func (r redisCacheStandalone) ZAdd(ctx context.Context, key string, member interface{}, score float64) error {
	logger.Debugf("ZAdd key: %s", key)
	data, err := sonic.Marshal(member)
	if err != nil {
		return err
	}

	result := r.client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: data,
	})
	return result.Err()
}

func (r redisCacheStandalone) ZPopMin(ctx context.Context, key string, count int) ([]string, error) {
	logger.Debugf("ZPopMin key: %s", key)
	result := r.client.ZPopMin(ctx, key, int64(count))
	if err := result.Err(); err != nil {
		return nil, err
	}

	// Convert the result to a slice of strings
	var poppedMembers []string
	for _, ele := range result.Val() {
		poppedMembers = append(poppedMembers, ele.Member.(string))
	}
	return poppedMembers, nil
}

func (r redisCacheStandalone) ZRCard(ctx context.Context, key string) (int64, error) {
	logger.Debugf("ZRCard key: %s", key)
	count, err := r.client.ZCard(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r redisCacheStandalone) ZRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	logger.Debugf("ZRange key: %s", key)
	result := r.client.ZRange(ctx, key, start, stop)

	var members []string
	for _, ele := range result.Val() {
		trimmed := strings.Trim(ele, `"`) // Remove surrounding quotes
		members = append(members, trimmed)
	}
	return members, result.Err()
}

func (r redisCacheStandalone) Expire(ctx context.Context, key string, ttl time.Duration) error {
	logger.Debugf("Expire key: %s", key)
	result := r.client.Expire(ctx, key, ttl)
	return result.Err()
}

func (r redisCacheStandalone) ZScore(ctx context.Context, keySet, member string) (float64, error) {
	logger.Debugf("ZScore key: %s", keySet)
	result := r.client.ZScore(ctx, keySet, member)
	return result.Val(), result.Err()
}

func (r redisCacheStandalone) Ping(ctx context.Context) error {
	logger.Debugf("Ping redis")
	result := r.client.Ping(ctx)
	return result.Err()
}

func (r redisCacheStandalone) Close() error {
	logger.Infof("Closing redis connection host: %s", r.client.Options().Addr)
	return r.client.Close()
}
