package amount

import (
	"github.com/program-world-labs/DDDGo/internal/domain/user/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	base_repository "github.com/program-world-labs/DDDGo/internal/infra/repository"
)

var _ repository.AmountRepository = (*RepoImpl)(nil)

type RepoImpl struct {
	base_repository.CRUDImpl
}

func NewRepoImpl(db datasource.IRelationDataSource, redis datasource.ICacheDataSource, cache datasource.ICacheDataSource) *RepoImpl {
	dtoAmount := &dto.Amount{}

	return &RepoImpl{CRUDImpl: *base_repository.NewCRUDImpl(db, redis, cache, dtoAmount)}
}
