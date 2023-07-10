package domain

import (
	"context"
)

type ICRUDRepository interface {
	GetByID(ctx context.Context, e IEntity) (IEntity, error)
	GetAll(ctx context.Context, e IEntity, sq *SearchQuery) ([]IEntity, error)
	Create(ctx context.Context, e IEntity) (IEntity, error)
	Update(ctx context.Context, e IEntity) (IEntity, error)
	UpdateWithFields(ctx context.Context, e IEntity, keys []string) (IEntity, error)
	Delete(ctx context.Context, e IEntity) error

	CreateTx(context.Context, IEntity, ITransactionEvent) (IEntity, error)
	UpdateTx(context.Context, IEntity, ITransactionEvent) (IEntity, error)
	UpdateWithFieldsTx(context.Context, IEntity, []string, ITransactionEvent) error
	DeleteTx(context.Context, IEntity, ITransactionEvent) error
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
