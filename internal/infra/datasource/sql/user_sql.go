package sql

import (
	entity "github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"gorm.io/gorm"
)

var _ datasource.IDataSource[*entity.User] = (*UserDatasourceImpl)(nil)

// UserDatasourceImpl -.
type UserDatasourceImpl struct {
	CRUDDatasourceImpl[*entity.User]
}

// NewUserDatasourceImpl -.
func NewUserDatasourceImpl(db *gorm.DB) *UserDatasourceImpl {
	return &UserDatasourceImpl{CRUDDatasourceImpl: *NewCRUDDatasourceImpl(db)}
}
