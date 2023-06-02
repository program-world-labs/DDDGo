package sql

import (
	"gorm.io/gorm"

	entity "github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
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
