package services

import (
	"context"
	"fmt"
	"time"

	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"github.com/redis/go-redis/v9"
)

// RedisService handles basic Redis operations
type RedisService struct {
	client *redis.Client
	ttl    time.Duration
}

// NewRedisService creates a new Redis service
func NewRedisService(cfg config.RedisConfig) (*RedisService, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})

	// Test the connection
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisService{
		client: rdb,
		ttl:    24 * time.Hour,
	}, nil
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

// Close closes the Redis connection
func (r *RedisService) Close() error {
	return r.client.Close()
}
