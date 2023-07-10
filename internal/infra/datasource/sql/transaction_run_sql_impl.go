package sql

import (
	"context"

	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/pkg/pwsql"
)

var _ datasource.ITransactionRun = (*TransactionDataSourceImpl)(nil)

// TransactionDataSourceImpl -.
type TransactionDataSourceImpl struct {
	DB *gorm.DB
}

// NewTransactionDataSourceImpl -.
func NewTransactionRunDataSourceImpl(db pwsql.ISQLGorm) *TransactionDataSourceImpl {
	return &TransactionDataSourceImpl{DB: db.GetDB()}
}

// RunTransaction -.
func (r *TransactionDataSourceImpl) RunTransaction(ctx context.Context, txFunc domain.TransactionEventFunc) error {
	// 創建一個新的Runner
	runner := func(tx *gorm.DB) error {
		// 創建一個新的Transaction
		txImpl := NewTransactionEventDataSourceImpl(tx)
		// 傳遞ctx到Runner
		return txFunc(ctx, txImpl)
	}
	// 傳遞ctx到RunTransaction
	return r.DB.Transaction(runner)
}
