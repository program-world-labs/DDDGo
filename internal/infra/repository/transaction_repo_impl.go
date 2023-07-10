package repository

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
)

var _ domain.ITransactionRepo = (*TransactionRunRepoImpl)(nil)

type TransactionRunRepoImpl struct {
	tr datasource.ITransactionRun
}

// NewTransactionRunRepoImpl -.
func NewTransactionRunRepoImpl(tr datasource.ITransactionRun) *TransactionRunRepoImpl {
	return &TransactionRunRepoImpl{tr: tr}
}

// RunTransaction -.
func (r *TransactionRunRepoImpl) RunTransaction(ctx context.Context, f domain.TransactionEventFunc) error {
	return r.tr.RunTransaction(ctx, f)
}
