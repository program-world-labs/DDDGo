package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
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

func (r *BigCacheDataSourceImpl) cacheKey(model dto.IRepoEntity, sq ...*domain.SearchQuery) string {
	if len(sq) > 0 {
		return fmt.Sprintf("%s-%s-%s", model.TableName(), model.GetID(), sq[0].GetKey())
	}

	return fmt.Sprintf("%s-%s", model.TableName(), model.GetID())
}

// Get -.
func (r *BigCacheDataSourceImpl) Get(_ context.Context, model dto.IRepoEntity, _ ...time.Duration) (dto.IRepoEntity, error) {
	data, err := r.Cache.Get(r.cacheKey(model))
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCacheGet, err)
	}

	err = json.Unmarshal(data, &model)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCacheGet, err)
	}

	return model, nil
}

// Set -.
func (r *BigCacheDataSourceImpl) Set(_ context.Context, model dto.IRepoEntity, _ ...time.Duration) (dto.IRepoEntity, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCacheSet, err)
	}

	err = r.Cache.Set(r.cacheKey(model), data)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCacheSet, err)
	}

	return model, nil
}

// Delete -.
func (r *BigCacheDataSourceImpl) Delete(_ context.Context, model dto.IRepoEntity) error {
	err := r.Cache.Delete(r.cacheKey(model))
	if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return domainerrors.Wrap(ErrorCodeCacheDelete, err)
	}

	return nil
}

// GetListItem -.
func (r *BigCacheDataSourceImpl) GetListItem(_ context.Context, model dto.IRepoEntity, sq *domain.SearchQuery, _ ...time.Duration) (*dto.List, error) {
	data, err := r.Cache.Get(r.cacheKey(model, sq))
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCacheGet, err)
	}

	var result *dto.List
	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCacheGet, err)
	}

	return result, nil
}

// SetListItem -.
func (r *BigCacheDataSourceImpl) SetListItem(_ context.Context, model []dto.IRepoEntity, sq *domain.SearchQuery, count int64, _ ...time.Duration) error {
	data, err := json.Marshal(model)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeCacheSet, err)
	}

	var info []map[string]interface{}
	err = json.Unmarshal(data, &info)

	if err != nil {
		return domainerrors.Wrap(ErrorCodeCacheSet, err)
	}

	var domainList = map[string]interface{}{
		"data":   info,
		"total":  count,
		"limit":  sq.Page.Limit,
		"offset": sq.Page.Offset,
	}

	result, err := json.Marshal(domainList)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeCacheSet, err)
	}

	err = r.Cache.Set(r.cacheKey(model[0], sq), result)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeCacheSet, err)
	}

	return nil
}

// DeleteListItem -.
func (r *BigCacheDataSourceImpl) DeleteListItem(_ context.Context, model dto.IRepoEntity, sq *domain.SearchQuery) error {
	err := r.Cache.Delete(r.cacheKey(model, sq))
	if err != nil {
		return domainerrors.Wrap(ErrorCodeCacheDelete, err)
	}

	return nil
}
