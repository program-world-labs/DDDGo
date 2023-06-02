package datasource

import (
	"context"

	entity "github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

type IEntityMethod interface {
	GetID() string
	SetID(string) error
	TableName() string
}
type IEntity interface {
	*entity.User
	IEntityMethod
}

type IDataSource[T IEntity] interface {
	GetByID(ctx context.Context, e T) (T, error)
	Create(ctx context.Context, e T) (T, error)
	Update(ctx context.Context, e T) (T, error)
	Delete(ctx context.Context, e T) error
}

type ICacheDataSource interface {
	Get(ctx context.Context, e IEntityMethod) (IEntityMethod, error)
	Set(ctx context.Context, e IEntityMethod) (IEntityMethod, error)
	Delete(ctx context.Context, e IEntityMethod) error
}
