package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/allegro/bigcache/v3"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
)

var _ datasource.ICacheDataSource = (*BigCacheDataSourceImpl)(nil)

// BigCacheDataSourceImpl -.
type BigCacheDataSourceImpl struct {
	Cache *bigcache.BigCache
}

// NewBigCacheDataSourceImpl -.
func NewBigCacheDataSourceImp(cache *bigcache.BigCache) *BigCacheDataSourceImpl {
	return &BigCacheDataSourceImpl{Cache: cache}
}

func (r *BigCacheDataSourceImpl) cacheKey(model domain.IEntity) string {
	return fmt.Sprintf("%s-%s", model.TableName(), model.GetID())
}

// Get -.
func (r *BigCacheDataSourceImpl) Get(_ context.Context, model domain.IEntity) (domain.IEntity, error) {
	data, err := r.Cache.Get(r.cacheKey(model))
	if err != nil {
		return nil, fmt.Errorf("BigCacheDataSourceImpl - Get - r.Cache.Get: %w", err)
	}

	err = json.Unmarshal(data, &model)
	if err != nil {
		return nil, fmt.Errorf("BigCacheDataSourceImpl - Get - json.Unmarshal: %w", err)
	}

	return model, nil
}

// Set -.
func (r *BigCacheDataSourceImpl) Set(_ context.Context, model domain.IEntity) (domain.IEntity, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, fmt.Errorf("BigCacheDataSourceImpl - Set - json.Marshal: %w", err)
	}

	err = r.Cache.Set(r.cacheKey(model), data)
	if err != nil {
		return nil, fmt.Errorf("BigCacheDataSourceImpl - Set - r.Cache.Set: %w", err)
	}

	return model, nil
}

// Delete -.
func (r *BigCacheDataSourceImpl) Delete(_ context.Context, model domain.IEntity) error {
	err := r.Cache.Delete(r.cacheKey(model))
	if err != nil {
		return fmt.Errorf("BigCacheDataSourceImpl - Delete - r.Cache.Delete: %w", err)
	}

	return nil
}
