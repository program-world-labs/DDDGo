package sql

import (
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
)

var _ domain.ITransactionEvent = (*TransactionEventDataSourceImpl)(nil)

type TransactionEventDataSourceImpl struct {
	DB *gorm.DB
}

func NewTransactionEventDataSourceImpl(db *gorm.DB) *TransactionEventDataSourceImpl {
	return &TransactionEventDataSourceImpl{DB: db}
}

func (r *TransactionEventDataSourceImpl) GetTx() interface{} {
	return r.DB
}
