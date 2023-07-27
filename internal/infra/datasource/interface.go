package datasource

import (
	"context"
	"time"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
)

type IDataSource interface {
	Create(context.Context, dto.IRepoEntity) (dto.IRepoEntity, error)
	Delete(context.Context, dto.IRepoEntity) (dto.IRepoEntity, error)
	Update(context.Context, dto.IRepoEntity) (dto.IRepoEntity, error)
	UpdateWithFields(context.Context, dto.IRepoEntity, []string) (dto.IRepoEntity, error)
	GetByID(context.Context, dto.IRepoEntity) (dto.IRepoEntity, error)
	GetAll(context.Context, *domain.SearchQuery, dto.IRepoEntity) (*dto.List, error)

	CreateTx(context.Context, dto.IRepoEntity, domain.ITransactionEvent) (dto.IRepoEntity, error)
	DeleteTx(context.Context, dto.IRepoEntity, domain.ITransactionEvent) (dto.IRepoEntity, error)
	UpdateTx(context.Context, dto.IRepoEntity, domain.ITransactionEvent) (dto.IRepoEntity, error)
	UpdateWithFieldsTx(context.Context, dto.IRepoEntity, []string, domain.ITransactionEvent) (dto.IRepoEntity, error)
}

type IAssociationDataSource interface {
	AppendAssociation(context.Context, string, dto.IRepoEntity, []dto.IRepoEntity) error
	ReplaceAssociation(context.Context, string, dto.IRepoEntity, []dto.IRepoEntity) error
	RemoveAssociation(context.Context, string, dto.IRepoEntity, []dto.IRepoEntity) error
	GetAssociationCount(context.Context, string, dto.IRepoEntity) (int64, error)

	AppendAssociationTx(context.Context, string, dto.IRepoEntity, []dto.IRepoEntity, domain.ITransactionEvent) error
	ReplaceAssociationTx(context.Context, string, dto.IRepoEntity, []dto.IRepoEntity, domain.ITransactionEvent) error
	RemoveAssociationTx(context.Context, string, dto.IRepoEntity, []dto.IRepoEntity, domain.ITransactionEvent) error
}

type IRelationDataSource interface {
	IDataSource
	IAssociationDataSource
}

type ICacheDataSource interface {
	Get(ctx context.Context, e dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error)
	Set(ctx context.Context, e dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error)
	Delete(ctx context.Context, e dto.IRepoEntity) error
	DeleteWithKey(ctx context.Context, key string) error
	GetListKeys(ctx context.Context, e dto.IRepoEntity) ([]string, error)
	GetListItem(ctx context.Context, e dto.IRepoEntity, sq *domain.SearchQuery, ttl ...time.Duration) (*dto.List, error)
	DeleteListKeys(ctx context.Context, e dto.IRepoEntity) error
	SetListItem(ctx context.Context, e []dto.IRepoEntity, sq *domain.SearchQuery, count int64, ttl ...time.Duration) error
}

type ITransactionRun interface {
	RunTransaction(context.Context, domain.TransactionEventFunc) error
}
