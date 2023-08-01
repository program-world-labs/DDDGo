package nosqlfs

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	firestoredb "github.com/program-world-labs/DDDGo/pkg/pwsql/nosql/firestoreDB"
)

var _ datasource.ITransactionRun = (*TransactionDataSourceImpl)(nil)

// TransactionDataSourceImpl -.
type TransactionDataSourceImpl struct {
	client *firestore.Client
}

// NewTransactionDataSourceImpl -.
func NewTransactionRunDataSourceImpl(db *firestoredb.Firestore) *TransactionDataSourceImpl {
	return &TransactionDataSourceImpl{client: db.GetClient()}
}

// RunTransaction -.
func (r *TransactionDataSourceImpl) RunTransaction(ctx context.Context, txFunc domain.TransactionEventFunc) error {
	// 創建一個新的Runner
	runner := func(ctx context.Context, tx *firestore.Transaction) error {
		// 創建一個新的Transaction
		txImpl := NewTransactionEventDataSourceImpl(tx)
		// 傳遞ctx到Runner
		return txFunc(ctx, txImpl)
	}
	// 傳遞ctx到RunTransaction
	return r.client.RunTransaction(ctx, runner)
}
