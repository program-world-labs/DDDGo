package repository

import (
	"github.com/program-world-labs/DDDGo/internal/domain/user/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
)

var _ repository.UserRepository = (*UserRepoImpl)(nil)

type UserRepoImpl struct {
	CRUDImpl
}

func NewUserRepoImpl(db datasource.IDataSource, redis datasource.ICacheDataSource, cache datasource.ICacheDataSource) *UserRepoImpl {
	return &UserRepoImpl{CRUDImpl: *NewCRUDImpl(db, redis, cache)}
}
