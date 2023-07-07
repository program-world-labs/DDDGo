package datasource

import (
	"context"
	"time"

	"github.com/program-world-labs/DDDGo/internal/infra/dto"
)

type IDataSource interface {
	GetByID(ctx context.Context, e dto.IRepoEntity) (dto.IRepoEntity, error)
	Create(ctx context.Context, e dto.IRepoEntity) (dto.IRepoEntity, error)
	Update(ctx context.Context, e dto.IRepoEntity) (dto.IRepoEntity, error)
	UpdateWithFields(ctx context.Context, e dto.IRepoEntity, m []string) (dto.IRepoEntity, error)
	Delete(ctx context.Context, e dto.IRepoEntity) error
}

type ICacheDataSource interface {
	Get(ctx context.Context, e dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error)
	Set(ctx context.Context, e dto.IRepoEntity, ttl ...time.Duration) (dto.IRepoEntity, error)
	Delete(ctx context.Context, e dto.IRepoEntity) error
}
