package user

import (
	"github.com/program-world-labs/DDDGo/internal/domain/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	base_repository "github.com/program-world-labs/DDDGo/internal/infra/repository"
)

var _ repository.UserRepository = (*RepoImpl)(nil)

type RepoImpl struct {
	base_repository.CRUDImpl
}

func NewRepoImpl(db datasource.IDataSource, redis datasource.ICacheDataSource, cache datasource.ICacheDataSource) *RepoImpl {
	dtoUser := &dto.User{}

	return &RepoImpl{CRUDImpl: *base_repository.NewCRUDImpl(db, redis, cache, dtoUser)}
}
