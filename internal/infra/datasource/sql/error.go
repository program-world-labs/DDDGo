package sql

import (
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeSQLCreate = domainerrors.ErrorCodeDatasourceSQL + iota + 1
	ErrorCodeSQLDelete
	ErrorCodeSQLUpdate
	ErrorCodeSQLUpdateWithFields
	ErrorCodeSQLGet
	ErrorCodeSQLGetAll
	ErrorCodeSQLCreateTx
	ErrorCodeSQLDeleteTx
	ErrorCodeSQLUpdateTx
	ErrorCodeSQLUpdateWithFieldsTx
	ErrorCodeSQLCast
	ErrorCodeSQLAppendAssociation
	ErrorCodeSQLReplaceAssociation
	ErrorCodeSQLRemoveAssociation
	ErrorCodeSQLGetAssociationCount
	ErrorCodeSQLAppendAssociationTx
	ErrorCodeSQLReplaceAssociationTx
	ErrorCodeSQLRemoveAssociationTx
)

var (
	ErrCastToEntityFailed = errors.New("cast to entity failed")
)
