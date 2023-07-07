package repository

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

type CacheUpdateImpl struct {
	RemoteCache datasource.ICacheDataSource
	Cache       datasource.ICacheDataSource
	DTOEntity   entity.IEntity
}

// NewCacheUpdateImpl -.
func NewCacheUpdateImpl(remoteCache datasource.ICacheDataSource, cache datasource.ICacheDataSource, data entity.IEntity) *CacheUpdateImpl {
	return &CacheUpdateImpl{RemoteCache: remoteCache, Cache: cache, DTOEntity: data}
}

// Save -.
func (r *CacheUpdateImpl) Save(ctx context.Context, e domain.IEntity) error {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return NewDatasourceError(err)
	}
	// 將資料寫入Redis
	_, err = r.RemoteCache.Set(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}
	// 將資料寫入Local Cache
	_, err = r.Cache.Set(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}

	return nil
}

// Delete -.
func (r *CacheUpdateImpl) Delete(ctx context.Context, e domain.IEntity) error {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return err
	}

	// 將資料從Redis刪除
	err = r.RemoteCache.Delete(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}
	// 將資料從Local Cache刪除
	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}

	return nil
}
