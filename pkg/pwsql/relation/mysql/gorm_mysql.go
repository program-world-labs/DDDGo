package mysql

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

// MySQL -.
type MySQL struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	db *gorm.DB
}

// New -.
func New(dsn string, opts ...Option) (*MySQL, error) {
	my := &MySQL{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(my)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = db.Use(tracing.NewPlugin(tracing.WithoutMetrics())); err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(my.maxPoolSize)
	sqlDB.SetMaxOpenConns(my.maxPoolSize)

	for my.connAttempts > 0 {
		err = sqlDB.PingContext(context.Background())
		if err == nil {
			break
		}

		log.Printf("MySQL is trying to connect, attempts left: %d", my.connAttempts)

		time.Sleep(my.connTimeout)

		my.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - NewMySQL - connAttempts == 0: %w", err)
	}

	my.db = db

	return my, nil
}

func (p *MySQL) GetDB() *gorm.DB {
	return p.db
}

// Close -.
func (p *MySQL) Close() error {
	if p.db != nil {
		sqlDB, err := p.db.DB()
		if err != nil {
			log.Printf("failed to get sql.DB: %v", err)

			return err
		}

		sqlDB.Close()
	}

	return nil
}
