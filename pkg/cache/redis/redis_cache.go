package redis_cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	_defaultMaxRetries = 3
	_defaultRetryDelay = time.Second
)

// Redis -.
type Redis struct {
	maxRetries int
	retryDelay time.Duration

	Client *redis.Client
}

// New -.
func New(dsn string, opts ...Option) (*Redis, error) {
	rd := &Redis{
		maxRetries: _defaultMaxRetries,
		retryDelay: _defaultRetryDelay,
	}

	// Custom options
	for _, opt := range opts {
		opt(rd)
	}

	options, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, fmt.Errorf("redis - NewRedis - ParseURL: %w", err)
	}

	client := redis.NewClient(options)

	for i := 0; i < rd.maxRetries; i++ {
		err := client.Ping(context.Background()).Err()
		if err == nil {
			break
		}

		if i == rd.maxRetries-1 {
			return nil, fmt.Errorf("redis - NewRedis - maxRetries == 0: %w", err)
		}

		time.Sleep(rd.retryDelay)
	}

	rd.Client = client

	return rd, nil
}

// Close -.
func (r *Redis) Close() {
	if r.Client != nil {
		err := r.Client.Close()
		if err != nil {
			log.Printf("failed to close redis client: %v", err)
			return
		}
	}
}
