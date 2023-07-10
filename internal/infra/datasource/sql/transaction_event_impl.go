package sql

import (
	"github.com/program-world-labs/DDDGo/internal/domain"
	"gorm.io/gorm"
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
