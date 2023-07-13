package sql

import (
	"errors"
	"fmt"

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

func NewCreateError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLCreate), err.Error())
}

func NewDeleteError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLDelete), err.Error())
}

func NewUpdateError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLUpdate), err.Error())
}

func NewUpdateWithFieldsError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLUpdateWithFields), err.Error())
}

func NewCreateTxError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLCreateTx), err.Error())
}

func NewDeleteTxError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLDeleteTx), err.Error())
}

func NewUpdateTxError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLUpdateTx), err.Error())
}

func NewUpdateWithFieldsTxError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLUpdateWithFieldsTx), err.Error())
}

func NewGetError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLGet), err.Error())
}

func NewGetAllError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLGetAll), err.Error())
}

func NewCastError(err error) *domainerrors.ErrorInfo {
	errCast := domainerrors.New(fmt.Sprint(ErrorCodeSQLCast), err.Error())

	return errCast
}

func NewAppendAssociationError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLAppendAssociation), err.Error())
}

func NewReplaceAssociationError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLReplaceAssociation), err.Error())
}

func NewRemoveAssociationError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLRemoveAssociation), err.Error())
}

func NewGetAssociationCountError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeSQLGetAssociationCount), err.Error())
}
