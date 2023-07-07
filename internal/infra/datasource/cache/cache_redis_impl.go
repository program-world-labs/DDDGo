package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/dtm-labs/rockscache"

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

func (r *RedisCacheDataSourceImpl) redisKey(model dto.IRepoEntity) string {
	return fmt.Sprintf("%s-%s", model.TableName(), model.GetID())
}

// GetByID -.
func (r *RedisCacheDataSourceImpl) Get(ctx context.Context, model dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error) {
	var t time.Duration
	if len(ttl) > 0 {
		t = ttl[0]
	}

	v, err := r.Client.Fetch2(ctx, r.redisKey(model), t, func() (string, error) {
		data, err := r.SQLDataSource.GetByID(ctx, model)
		if err != nil {
			return "", err
		}

		return data.ToJSON()
	})
	if err != nil {
		return nil, NewGetError(err)
	}

	err = model.DecodeJSON(v)
	if err != nil {
		return nil, NewGetError(err)
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
		return NewDeleteError(err)
	}

	return nil
}
