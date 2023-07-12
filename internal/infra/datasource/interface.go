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
	UpdateWithFields(context.Context, dto.IRepoEntity, []string) error
	GetByID(context.Context, dto.IRepoEntity) (map[string]interface{}, error)
	GetAll(context.Context, *domain.SearchQuery, dto.IRepoEntity) (map[string]interface{}, error)

	CreateTx(context.Context, dto.IRepoEntity, domain.ITransactionEvent) (dto.IRepoEntity, error)
	DeleteTx(context.Context, dto.IRepoEntity, domain.ITransactionEvent) error
	UpdateTx(context.Context, dto.IRepoEntity, domain.ITransactionEvent) (dto.IRepoEntity, error)
	UpdateWithFieldsTx(context.Context, dto.IRepoEntity, []string, domain.ITransactionEvent) error
}

type ICacheDataSource interface {
	Get(ctx context.Context, e dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error)
	Set(ctx context.Context, e dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error)
	Delete(ctx context.Context, e dto.IRepoEntity) error
	GetListItem(ctx context.Context, e dto.IRepoEntity, sq *domain.SearchQuery, ttl ...time.Duration) (map[string]interface{}, error)
	SetListItem(ctx context.Context, e []dto.IRepoEntity, sq *domain.SearchQuery, count int64, ttl ...time.Duration) error
	DeleteListItem(ctx context.Context, e dto.IRepoEntity, sq *domain.SearchQuery) error
}

type ITransactionRun interface {
	RunTransaction(context.Context, domain.TransactionEventFunc) error
}
