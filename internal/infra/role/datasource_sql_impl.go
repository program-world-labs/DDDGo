package role

import (
	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource/sql"
	"github.com/program-world-labs/DDDGo/pkg/pwsql"
)

var _ datasource.IDataSource = (*DatasourceImpl)(nil)

// DatasourceImpl -.
type DatasourceImpl struct {
	sql.CRUDDatasourceImpl
}

// NewDatasourceImpl -.
func NewDatasourceImpl(db pwsql.ISQLGorm) *DatasourceImpl {
	return &DatasourceImpl{CRUDDatasourceImpl: *sql.NewCRUDDatasourceImpl(db)}
}
