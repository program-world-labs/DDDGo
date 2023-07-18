package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dtm-labs/rockscache"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
)

var _ datasource.ICacheDataSource = (*RedisCacheDataSourceImpl)(nil)

// RedisCacheDataSourceImpl -.
type RedisCacheDataSourceImpl struct {
	Client        *rockscache.Client
	SQLDataSource datasource.IDataSource
}

// NewRedisCacheDataSourceImpl -.
func NewRedisCacheDataSourceImpl(client *rockscache.Client, sqlDatasource datasource.IDataSource) *RedisCacheDataSourceImpl {
	return &RedisCacheDataSourceImpl{Client: client, SQLDataSource: sqlDatasource}
}

func (r *RedisCacheDataSourceImpl) redisKey(model dto.IRepoEntity, sq ...*domain.SearchQuery) string {
	if len(sq) > 0 {
		return fmt.Sprintf("%s-%s-%s", model.TableName(), model.GetID(), sq[0].GetKey())
	}

	return fmt.Sprintf("%s-%s", model.TableName(), model.GetID())
}

// GetByID -.
func (r *RedisCacheDataSourceImpl) Get(ctx context.Context, model dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error) {
	// 預設10秒
	const defaultTTL = 60 * time.Second

	var t time.Duration

	if len(ttl) > 0 {
		t = ttl[0]
	} else {
		t = defaultTTL
	}

	v, err := r.Client.Fetch2(ctx, r.redisKey(model), t, func() (string, error) {
		data, err := r.SQLDataSource.GetByID(ctx, model)
		if err != nil {
			return "", err
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			return "", err
		}

		return string(jsonData), nil
	})
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCacheGet, err)
	}

	err = json.Unmarshal([]byte(v), &model)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCacheGet, err)
	}

	return model, nil
}

// Set -.
func (r *RedisCacheDataSourceImpl) Set(_ context.Context, model dto.IRepoEntity, _ ...time.Duration) (dto.IRepoEntity, error) {
	// data, err := json.Marshal(model)
	// if err != nil {
	// 	return nil, fmt.Errorf("RedisCacheDataSourceImpl - Create - json.Marshal: %w", err)
	// }

	// err = r.Client.Set(ctx, r.redisKey(model), data, 0).Err()
	// if err != nil {
	// 	return nil, fmt.Errorf("RedisCacheDataSourceImpl - Create - r.Client.Set: %w", err)
	// }

	return model, nil
}

// Delete -.
func (r *RedisCacheDataSourceImpl) Delete(ctx context.Context, model dto.IRepoEntity) error {
	err := r.Client.TagAsDeleted2(ctx, r.redisKey(model))
	if err != nil {
		return domainerrors.Wrap(ErrorCodeCacheDelete, err)
	}

	return nil
}

// GetListItem -.
func (r *RedisCacheDataSourceImpl) GetListItem(ctx context.Context, model dto.IRepoEntity, sq *domain.SearchQuery, ttl ...time.Duration) (*dto.List, error) {
	const defaultTTL = 60 * time.Second

	var t time.Duration

	if len(ttl) > 0 {
		t = ttl[0]
	} else {
		t = defaultTTL
	}

	v, err := r.Client.Fetch(r.redisKey(model, sq), t, func() (string, error) {
		data, err := r.SQLDataSource.GetAll(ctx, sq, model)
		if err != nil {
			return "", err
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			return "", err
		}

		return string(jsonData), nil
	})
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCacheGet, err)
	}

	var value *dto.List
	err = json.Unmarshal([]byte(v), &value)

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCacheGet, err)
	}

	return value, nil
}

// SetListItem -.
func (r *RedisCacheDataSourceImpl) SetListItem(_ context.Context, _ []dto.IRepoEntity, _ *domain.SearchQuery, _ int64, _ ...time.Duration) error {
	return nil
}

// DeleteListItem -.
func (r *RedisCacheDataSourceImpl) DeleteListItem(ctx context.Context, model dto.IRepoEntity, sq *domain.SearchQuery) error {
	err := r.Client.TagAsDeleted2(ctx, r.redisKey(model, sq))
	if err != nil {
		return domainerrors.Wrap(ErrorCodeCacheDelete, err)
	}

	return nil
}
