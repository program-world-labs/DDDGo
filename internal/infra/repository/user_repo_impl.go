package repository

import (
	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/domain/user/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
)

var _ repository.UserRepository = (*UserRepoImpl)(nil)

type UserRepoImpl struct {
	CRUDImpl[*entity.User]
}

func NewUserRepoImpl(db datasource.DataSource[*entity.User], redis datasource.CacheDataSource[*entity.User], cache datasource.CacheDataSource[*entity.User]) *UserRepoImpl {
	return &UserRepoImpl{CRUDImpl: *NewCRUDImpl(db, redis, cache)}
}
