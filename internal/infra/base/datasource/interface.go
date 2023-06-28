package datasource

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

type IDataSource interface {
	GetByID(ctx context.Context, e entity.IEntity) (entity.IEntity, error)
	Create(ctx context.Context, e entity.IEntity) (entity.IEntity, error)
	Update(ctx context.Context, e entity.IEntity) (entity.IEntity, error)
	UpdateWithFields(ctx context.Context, e entity.IEntity, m []string) (entity.IEntity, error)
	Delete(ctx context.Context, e entity.IEntity) error
}

type ICacheDataSource interface {
	Get(ctx context.Context, e entity.IEntity) (entity.IEntity, error)
	Set(ctx context.Context, e entity.IEntity) (entity.IEntity, error)
	Delete(ctx context.Context, e entity.IEntity) error
}
