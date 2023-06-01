package cache

import (
	"context"
	"encoding/json"
	"fmt"

	entity "github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/redis/go-redis/v9"
)

var _ datasource.CacheDataSource[*entity.User] = (*RedisDataSourceImpl[*entity.User])(nil)

// RedisDataSourceImpl -.
type RedisDataSourceImpl[T datasource.Entity] struct {
	Client *redis.Client
}

// NewRedisDataSourceImpl -.
func NewRedisDataSourceImpl[T datasource.Entity](client *redis.Client) *RedisDataSourceImpl[T] {
	return &RedisDataSourceImpl[T]{Client: client}
}

func (r *RedisDataSourceImpl[T]) redisKey(model T) string {
	return fmt.Sprintf("%s-%s", model.TableName(), model.GetID())
}

// GetByID -.
func (r *RedisDataSourceImpl[T]) Get(ctx context.Context, model T) (T, error) {
	data, err := r.Client.Get(ctx, r.redisKey(model)).Bytes()
	if err != nil {
		return nil, fmt.Errorf("RedisDataSourceImpl - GetByID - r.Client.Get: %w", err)
	}

	err = json.Unmarshal(data, &model)
	if err != nil {
		return nil, fmt.Errorf("RedisDataSourceImpl - GetByID - json.Unmarshal: %w", err)
	}

	return model, nil
}

// Set -.
func (r *RedisDataSourceImpl[T]) Set(ctx context.Context, model T) (T, error) {
	data, err := json.Marshal(model)
	if err != nil {
		return nil, fmt.Errorf("RedisDataSourceImpl - Create - json.Marshal: %w", err)
	}

	err = r.Client.Set(ctx, r.redisKey(model), data, 0).Err()
	if err != nil {
		return nil, fmt.Errorf("RedisDataSourceImpl - Create - r.Client.Set: %w", err)
	}

	return model, nil
}

// Delete -.
func (r *RedisDataSourceImpl[T]) Delete(ctx context.Context, model T) error {
	err := r.Client.Del(ctx, r.redisKey(model)).Err()
	if err != nil {
		return fmt.Errorf("RedisDataSourceImpl - Delete - r.Client.Del: %w", err)
	}

	return nil
}
