package repository

import (
	"context"
	"fmt"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

type CacheUpdateImpl struct {
	RemoteCache  datasource.ICacheDataSource
	Cache        datasource.ICacheDataSource
	entityEntity entity.IEntity
}

// NewCacheUpdateImpl -.
func NewCacheUpdateImpl(remoteCache datasource.ICacheDataSource, cache datasource.ICacheDataSource, data entity.IEntity) *CacheUpdateImpl {
	return &CacheUpdateImpl{RemoteCache: remoteCache, Cache: cache, entityEntity: data}
}

// Save -.
func (r *CacheUpdateImpl) Save(ctx context.Context, e domain.IEntity) error {
	info, err := r.entityEntity.Transform(e)
	if err != nil {
		return fmt.Errorf("CacheUpdateImpl - Save - r.entityEntity.Transform: %w", err)
	}
	// 將資料寫入Redis
	_, err = r.RemoteCache.Set(ctx, info)
	if err != nil {
		return fmt.Errorf("CacheUpdateImpl - Save - r.Redis.Save: %w", err)
	}
	// 將資料寫入Local Cache
	_, err = r.Cache.Set(ctx, info)
	if err != nil {
		return fmt.Errorf("CacheUpdateImpl - Save - r.Cache.Save: %w", err)
	}

	return nil
}

// Delete -.
func (r *CacheUpdateImpl) Delete(ctx context.Context, e domain.IEntity) error {
	info, err := r.entityEntity.Transform(e)
	if err != nil {
		return fmt.Errorf("CacheUpdateImpl - Delete - r.entityEntity.Transform: %w", err)
	}

	// 將資料從Redis刪除
	err = r.RemoteCache.Delete(ctx, info)
	if err != nil {
		return fmt.Errorf("CacheUpdateImpl - Delete - r.Redis.Delete: %w", err)
	}
	// 將資料從Local Cache刪除
	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return fmt.Errorf("CacheUpdateImpl - Delete - r.Cache.Delete: %w", err)
	}

	return nil
}
