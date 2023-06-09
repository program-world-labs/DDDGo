package sql

import (
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
)

var _ datasource.IDataSource = (*UserDatasourceImpl)(nil)

// UserDatasourceImpl -.
type UserDatasourceImpl struct {
	CRUDDatasourceImpl
}

// NewUserDatasourceImpl -.
func NewUserDatasourceImpl(db *gorm.DB) *UserDatasourceImpl {
	return &UserDatasourceImpl{CRUDDatasourceImpl: *NewCRUDDatasourceImpl(db)}
}
