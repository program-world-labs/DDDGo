package domain

import (
	"context"
)

type ICRUDRepository interface {
	GetByID(ctx context.Context, e IEntity) (IEntity, error)
	Create(ctx context.Context, e IEntity) (IEntity, error)
	Update(ctx context.Context, e IEntity) (IEntity, error)
	Delete(ctx context.Context, e IEntity) error
}

type ICacheUpdateRepository interface {
	Save(ctx context.Context, e IEntity) error
	Delete(ctx context.Context, e IEntity) error
}
