package datasource

import (
	"context"

	entity "gitlab.com/demojira/template.git/internal/domain/user/entity"
)

type Entity interface {
	*entity.User
	GetID() string
	SetID(string)
	TableName() string
}
type DataSource[T Entity] interface {
	GetByID(ctx context.Context, e T) (T, error)
	Create(ctx context.Context, e T) (T, error)
	Update(ctx context.Context, e T) (T, error)
	Delete(ctx context.Context, e T) error
}

type CacheDataSource[T Entity] interface {
	Get(ctx context.Context, e T) (T, error)
	Set(ctx context.Context, e T) (T, error)
	Delete(ctx context.Context, e T) error
}
