package sql

import (
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeSQLCreate = domainerrors.ErrorCodeInfraDatasource + domainerrors.ErrorCodeInfraDatasource + domainerrors.ErrorCodeInfraDatasourceSQL + iota
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
