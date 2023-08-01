package nosqlfs

import (
	"cloud.google.com/go/firestore"

	"github.com/program-world-labs/DDDGo/internal/domain"
)

var _ domain.ITransactionEvent = (*TransactionEventDataSourceImpl)(nil)

type TransactionEventDataSourceImpl struct {
	tx *firestore.Transaction
}

func NewTransactionEventDataSourceImpl(db *firestore.Transaction) *TransactionEventDataSourceImpl {
	return &TransactionEventDataSourceImpl{tx: db}
}

func (r *TransactionEventDataSourceImpl) GetTx() interface{} {
	return r.tx
}
