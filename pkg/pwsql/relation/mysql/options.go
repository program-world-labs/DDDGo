package mysql

import "time"

// Option -.
type Option func(*MySQL)

// MaxPoolSize -.
func MaxPoolSize(size int) Option {
	return func(c *MySQL) {
		c.maxPoolSize = size
	}
}

// ConnAttempts -.
func ConnAttempts(attempts int) Option {
	return func(c *MySQL) {
		c.connAttempts = attempts
	}
}

// ConnTimeout -.
func ConnTimeout(timeout time.Duration) Option {
	return func(c *MySQL) {
		c.connTimeout = timeout
	}
}
