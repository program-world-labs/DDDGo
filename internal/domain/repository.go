package domain

import (
	"context"

	"gitlab.com/demojira/template.git/internal/infra/datasource"
)

type CRUDRepository[T datasource.Entity] interface {
	GetByID(ctx context.Context, e T) (T, error)
	Create(ctx context.Context, e T) (T, error)
	Update(ctx context.Context, e T) (T, error)
	Delete(ctx context.Context, e T) error
}

type CacheUpdateRepository[T datasource.Entity] interface {
	Save(ctx context.Context, e T) error
	Delete(ctx context.Context, e T) error
}
