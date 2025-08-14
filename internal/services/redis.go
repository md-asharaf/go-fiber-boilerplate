package services

import (
	"context"
	"fmt"
	"time"

	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// RedisService handles Redis operations
type RedisService struct {
	client *redis.Client
	ttl    time.Duration
}

// NewRedisService creates a new Redis service
func NewRedisService(cfg config.RedisConfig) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0, // Default to DB 0
	})

	return &RedisService{
		client: rdb,
		ttl:    24 * time.Hour, // Default 24 hour TTL
	}
}

// Ping checks if Redis is connected
func (r *RedisService) Ping(ctx context.Context) error {
	_, err := r.client.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("Redis ping failed", zap.Error(err))
		return err
	}
	zap.L().Info("Redis connection established")
	return nil
}

// Set stores a key-value pair with TTL
func (r *RedisService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if expiration == 0 {
		expiration = r.ttl
	}
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value by key
func (r *RedisService) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

// Delete removes a key
func (r *RedisService) Delete(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

// Exists checks if a key exists
func (r *RedisService) Exists(ctx context.Context, key string) (bool, error) {
	count, err := r.client.Exists(ctx, key).Result()
	return count > 0, err
}

// SetNX sets a key only if it doesn't exist (atomic operation)
func (r *RedisService) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) (bool, error) {
	if expiration == 0 {
		expiration = r.ttl
	}
	return r.client.SetNX(ctx, key, value, expiration).Result()
}

// Increment increments a key's value
func (r *RedisService) Increment(ctx context.Context, key string) (int64, error) {
	return r.client.Incr(ctx, key).Result()
}

// IncrementWithExpire increments a key and sets expiration if key is new
func (r *RedisService) IncrementWithExpire(ctx context.Context, key string, expiration time.Duration) (int64, error) {
	pipe := r.client.Pipeline()
	incr := pipe.Incr(ctx, key)
	pipe.Expire(ctx, key, expiration)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return 0, err
	}

	return incr.Val(), nil
}

// HSet sets a hash field
func (r *RedisService) HSet(ctx context.Context, key string, field string, value interface{}) error {
	return r.client.HSet(ctx, key, field, value).Err()
}

// HGet gets a hash field
func (r *RedisService) HGet(ctx context.Context, key string, field string) (string, error) {
	return r.client.HGet(ctx, key, field).Result()
}

// HGetAll gets all hash fields
func (r *RedisService) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, key).Result()
}

// HDel deletes hash fields
func (r *RedisService) HDel(ctx context.Context, key string, fields ...string) error {
	return r.client.HDel(ctx, key, fields...).Err()
}

// Close closes the Redis connection
func (r *RedisService) Close() error {
	return r.client.Close()
}

// GetClient returns the Redis client for advanced operations
func (r *RedisService) GetClient() *redis.Client {
	return r.client
}
