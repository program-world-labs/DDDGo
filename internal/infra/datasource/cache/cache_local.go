package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/allegro/bigcache/v3"
	entity "github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
)

var _ datasource.CacheDataSource[*entity.User] = (*RedisDataSourceImpl[*entity.User])(nil)

// BigCacheDataSourceImpl -.
type BigCacheDataSourceImpl[T datasource.Entity] struct {
	Cache *bigcache.BigCache
}

// NewBigCacheDataSourceImpl -.
func NewBigCacheDataSourceImpl[T datasource.Entity](cache *bigcache.BigCache) (*BigCacheDataSourceImpl[T], error) {
	return &BigCacheDataSourceImpl[T]{Cache: cache}, nil
}

func (r *BigCacheDataSourceImpl[T]) cacheKey(model T) string {
	return fmt.Sprintf("%s-%s", model.TableName(), model.GetID())
}

// Get -.
func (r *BigCacheDataSourceImpl[T]) Get(ctx context.Context, model T) (T, error) {
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
func (r *BigCacheDataSourceImpl[T]) Set(ctx context.Context, model T) (T, error) {
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
func (r *BigCacheDataSourceImpl[T]) Delete(ctx context.Context, model T) error {
	err := r.Cache.Delete(r.cacheKey(model))
	if err != nil {
		return fmt.Errorf("BigCacheDataSourceImpl - Delete - r.Cache.Delete: %w", err)
	}

	return nil
}
