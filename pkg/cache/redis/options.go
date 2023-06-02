package redis_cache

import "time"

// Option -.
type Option func(*Redis)

// MaxRetries -.
func MaxRetries(attempts int) Option {
	return func(c *Redis) {
		c.maxRetries = attempts
	}
}

// RetryDelay -.
func RetryDelay(t time.Duration) Option {
	return func(c *Redis) {
		c.retryDelay = t
	}
}
