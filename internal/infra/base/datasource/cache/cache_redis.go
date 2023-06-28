package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

var _ datasource.ICacheDataSource = (*RedisCacheDataSourceImpl)(nil)

// RedisCacheDataSourceImpl -.
type RedisCacheDataSourceImpl struct {
	Client *redis.Client
}

// NewRedisCacheDataSourceImpl -.
func NewRedisCacheDataSourceImpl(client *redis.Client) *RedisCacheDataSourceImpl {
	return &RedisCacheDataSourceImpl{Client: client}
}

func (r *RedisCacheDataSourceImpl) redisKey(model entity.IEntity) string {
	return fmt.Sprintf("%s-%s", model.TableName(), model.GetID())
}

// GetByID -.
func (r *RedisCacheDataSourceImpl) Get(ctx context.Context, model entity.IEntity) (entity.IEntity, error) {
	data, err := r.Client.Get(ctx, r.redisKey(model)).Bytes()
	if err != nil {
		return nil, fmt.Errorf("RedisCacheDataSourceImpl - GetByID - r.Client.Get: %w", err)
	}

	err = json.Unmarshal(data, &model)
	if err != nil {
		return nil, fmt.Errorf("RedisCacheDataSourceImpl - GetByID - json.Unmarshal: %w", err)
	}

	return model, nil
}

// Set -.
func (r *RedisCacheDataSourceImpl) Set(ctx context.Context, model entity.IEntity) (entity.IEntity, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, fmt.Errorf("RedisCacheDataSourceImpl - Create - json.Marshal: %w", err)
	}

	err = r.Client.Set(ctx, r.redisKey(model), data, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("RedisCacheDataSourceImpl - Create - r.Client.Set: %w", err)
	}

	return model, nil
}

// Delete -.
func (r *RedisCacheDataSourceImpl) Delete(ctx context.Context, model entity.IEntity) error {
	err := r.Client.Del(ctx, r.redisKey(model)).Err()
	if err != nil {
		return fmt.Errorf("RedisCacheDataSourceImpl - Delete - r.Client.Del: %w", err)
	}

	return nil
}
