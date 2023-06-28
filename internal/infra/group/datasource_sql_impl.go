package group

import (
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource/sql"
)

var _ datasource.IDataSource = (*DatasourceImpl)(nil)

// DatasourceImpl -.
type DatasourceImpl struct {
	sql.CRUDDatasourceImpl
}

// NewDatasourceImpl -.
func NewDatasourceImpl(db *gorm.DB) *DatasourceImpl {
	return &DatasourceImpl{CRUDDatasourceImpl: *sql.NewCRUDDatasourceImpl(db)}
}
