package wallet

import (
	"github.com/program-world-labs/DDDGo/internal/domain/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	base_repository "github.com/program-world-labs/DDDGo/internal/infra/repository"
)

var _ repository.WalletBalanceRepository = (*RepoImpl)(nil)

type RepoImpl struct {
	base_repository.CRUDImpl
}

func NewRepoImpl(db datasource.IRelationDataSource, redis datasource.ICacheDataSource, cache datasource.ICacheDataSource) *RepoImpl {
	dtoWalletBalance := &dto.WalletBalance{}

	return &RepoImpl{CRUDImpl: *base_repository.NewCRUDImpl(db, redis, cache, dtoWalletBalance)}
}
