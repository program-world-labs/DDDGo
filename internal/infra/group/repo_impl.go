package group

import (
	"github.com/program-world-labs/DDDGo/internal/domain/user/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource"
	base_repository "github.com/program-world-labs/DDDGo/internal/infra/base/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
)

var _ repository.GroupRepository = (*RepoImpl)(nil)

type RepoImpl struct {
	base_repository.CRUDImpl
}

func NewRepoImpl(db datasource.IDataSource, redis datasource.ICacheDataSource, cache datasource.ICacheDataSource) *RepoImpl {
	dtoGroup := &dto.Group{}

	return &RepoImpl{CRUDImpl: *base_repository.NewCRUDImpl(db, redis, cache, dtoGroup)}
}
