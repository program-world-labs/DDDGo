package repository

import (
	"context"
	"fmt"

	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
)

type CacheUpdateImpl[T datasource.Entity] struct {
	Redis datasource.CacheDataSource[T]
	Cache datasource.CacheDataSource[T]
}

// NewCacheUpdateImpl -.
func NewCacheUpdateImpl[T datasource.Entity](redis datasource.CacheDataSource[T], cache datasource.CacheDataSource[T]) *CacheUpdateImpl[T] {
	return &CacheUpdateImpl[T]{Redis: redis, Cache: cache}
}

// Save -.
func (r *CacheUpdateImpl[T]) Save(ctx context.Context, e T) error {
	// 將資料寫入Redis
	_, err := r.Redis.Set(ctx, e)
	if err != nil {
		return fmt.Errorf("CacheUpdateImpl - Save - r.Redis.Save: %w", err)
	}
	// 將資料寫入Local Cache
	_, err = r.Cache.Set(ctx, e)
	if err != nil {
		return fmt.Errorf("CacheUpdateImpl - Save - r.Cache.Save: %w", err)
	}

	return nil
}

// Delete -.
func (r *CacheUpdateImpl[T]) Delete(ctx context.Context, e T) error {
	// 將資料從Redis刪除
	err := r.Redis.Delete(ctx, e)
	if err != nil {
		return fmt.Errorf("CacheUpdateImpl - Delete - r.Redis.Delete: %w", err)
	}
	// 將資料從Local Cache刪除
	err = r.Cache.Delete(ctx, e)
	if err != nil {
		return fmt.Errorf("CacheUpdateImpl - Delete - r.Cache.Delete: %w", err)
	}

	return nil
}
