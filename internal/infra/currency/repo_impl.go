package currency

import (
	"github.com/program-world-labs/DDDGo/internal/domain/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	base_repository "github.com/program-world-labs/DDDGo/internal/infra/repository"
)

var _ repository.CurrencyRepository = (*RepoImpl)(nil)

type RepoImpl struct {
	base_repository.CRUDImpl
}

func NewRepoImpl(db datasource.IRelationDataSource, redis datasource.ICacheDataSource, cache datasource.ICacheDataSource) *RepoImpl {
	dtoCurrency := &dto.Currency{}

	return &RepoImpl{CRUDImpl: *base_repository.NewCRUDImpl(db, redis, cache, dtoCurrency)}
}
