package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"

	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
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

func (r *BigCacheDataSourceImpl) cacheKey(model entity.IEntity) string {
	return fmt.Sprintf("%s-%s", model.TableName(), model.GetID())
}

// Get -.
func (r *BigCacheDataSourceImpl) Get(_ context.Context, model entity.IEntity, _ ...time.Duration) (entity.IEntity, error) {
	data, err := r.Cache.Get(r.cacheKey(model))
	if err != nil {
		return nil, NewGetError(err)
	}

	err = json.Unmarshal(data, &model)
	if err != nil {
		return nil, NewGetError(err)
	}

	return model, nil
}

// Set -.
func (r *BigCacheDataSourceImpl) Set(_ context.Context, model entity.IEntity, _ ...time.Duration) (entity.IEntity, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, NewSetError(err)
	}

	err = r.Cache.Set(r.cacheKey(model), data)
	if err != nil {
		return nil, NewSetError(err)
	}

	return model, nil
}

// Delete -.
func (r *BigCacheDataSourceImpl) Delete(_ context.Context, model entity.IEntity) error {
	err := r.Cache.Delete(r.cacheKey(model))
	if err != nil {
		return NewDeleteError(err)
	}

	return nil
}
