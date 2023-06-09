package datasource

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/domain"
)

type IDataSource interface {
	GetByID(ctx context.Context, e domain.IEntity) (domain.IEntity, error)
	Create(ctx context.Context, e domain.IEntity) (domain.IEntity, error)
	Update(ctx context.Context, e domain.IEntity) (domain.IEntity, error)
	UpdateWithFields(ctx context.Context, e domain.IEntity, m []string) (domain.IEntity, error)
	Delete(ctx context.Context, e domain.IEntity) error
}

type ICacheDataSource interface {
	Get(ctx context.Context, e domain.IEntity) (domain.IEntity, error)
	Set(ctx context.Context, e domain.IEntity) (domain.IEntity, error)
	Delete(ctx context.Context, e domain.IEntity) error
}
