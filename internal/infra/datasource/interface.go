package datasource

import (
	"context"
	"time"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
)

type IDataSource interface {
	Create(context.Context, dto.IRepoEntity) (dto.IRepoEntity, error)
	Delete(context.Context, dto.IRepoEntity) error
	Update(context.Context, dto.IRepoEntity) (dto.IRepoEntity, error)
	UpdateWithFields(context.Context, dto.IRepoEntity, []string) (dto.IRepoEntity, error)
	GetByID(context.Context, dto.IRepoEntity) (dto.IRepoEntity, error)
	GetAll(context.Context, dto.IRepoEntity, *domain.SearchQuery) ([]dto.IRepoEntity, error)

	CreateTx(context.Context, dto.IRepoEntity, domain.ITransactionEvent) (dto.IRepoEntity, error)
	DeleteTx(context.Context, dto.IRepoEntity, domain.ITransactionEvent) error
	UpdateTx(context.Context, dto.IRepoEntity, domain.ITransactionEvent) (dto.IRepoEntity, error)
	UpdateWithFieldsTx(context.Context, dto.IRepoEntity, []string, domain.ITransactionEvent) error
}

type ICacheDataSource interface {
	Get(ctx context.Context, e dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error)
	Set(ctx context.Context, e dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error)
	Delete(ctx context.Context, e dto.IRepoEntity) error
}

type ITransactionRun interface {
	RunTransaction(context.Context, domain.TransactionEventFunc) error
}
