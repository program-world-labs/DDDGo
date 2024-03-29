package domain

import (
	"context"
)

type ICRUDRepository interface {
	GetByID(ctx context.Context, e IEntity) (IEntity, error)
	GetAll(ctx context.Context, sq *SearchQuery, e IEntity) (*List, error)
	Create(ctx context.Context, e IEntity) (IEntity, error)
	Update(ctx context.Context, e IEntity) (IEntity, error)
	UpdateWithFields(ctx context.Context, e IEntity, keys []string) (IEntity, error)
	Delete(ctx context.Context, e IEntity) (IEntity, error)

	CreateTx(context.Context, IEntity, ITransactionEvent) (IEntity, error)
	UpdateTx(context.Context, IEntity, ITransactionEvent) (IEntity, error)
	UpdateWithFieldsTx(context.Context, IEntity, []string, ITransactionEvent) (IEntity, error)
	DeleteTx(context.Context, IEntity, ITransactionEvent) (IEntity, error)
}

type ICacheUpdateRepository interface {
	Save(ctx context.Context, e IEntity) error
	Delete(ctx context.Context, e IEntity) error
}

type TransactionEventFunc func(context.Context, ITransactionEvent) error

type ITransactionRepo interface {
	RunTransaction(ctx context.Context, f TransactionEventFunc) error
}

type ITransactionEvent interface {
	GetTx() interface{}
}
