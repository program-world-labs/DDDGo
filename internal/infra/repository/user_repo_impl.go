package repository

import (
	"gitlab.com/demojira/template.git/internal/domain/user/entity"
	"gitlab.com/demojira/template.git/internal/domain/user/repository"
	"gitlab.com/demojira/template.git/internal/infra/datasource"
)

var _ repository.UserRepository = (*UserRepoImpl)(nil)

type UserRepoImpl struct {
	CRUDImpl[*entity.User]
}

func NewUserRepoImpl(db datasource.DataSource[*entity.User], redis datasource.CacheDataSource[*entity.User], cache datasource.CacheDataSource[*entity.User]) *UserRepoImpl {
	return &UserRepoImpl{CRUDImpl: *NewCRUDImpl(db, redis, cache)}
}
