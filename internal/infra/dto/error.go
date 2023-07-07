package dto

import (
	"fmt"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeUserTransform = domainerrors.ErrorCodeDatasourceUserRepoDTO + iota + 1
	ErrorCodeUserToJSON
	ErrorCodeUserDecodeJSON
	ErrorCodeRoleTransform = domainerrors.ErrorCodeDatasourceRoleRepoDTO + iota + 1
	ErrorCodeRoleToJSON
	ErrorCodeRoleDecodeJSON
	ErrorCodeGroupTransform = domainerrors.ErrorCodeDatasourceGroupRepoDTO + iota + 1
	ErrorCodeGroupToJSON
	ErrorCodeGroupDecodeJSON
	ErrorCodeWalletTransform = domainerrors.ErrorCodeDatasourceWalletRepoDTO + iota + 1
	ErrorCodeWalletToJSON
	ErrorCodeWalletDecodeJSON
	ErrorCodeAmountTransform = domainerrors.ErrorCodeDatasourceAmountRepoDTO + iota + 1
	ErrorCodeAmountToJSON
	ErrorCodeAmountDecodeJSON
)

func NewUserTransformError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeUserTransform), err.Error())
}

func NewUserToJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeUserToJSON), err.Error())
}

func NewUserDecodeJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeUserDecodeJSON), err.Error())
}

func NewRoleTransformError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeRoleTransform), err.Error())
}

func NewRoleToJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeRoleToJSON), err.Error())
}

func NewRoleDecodeJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeRoleDecodeJSON), err.Error())
}
func NewGroupTransformError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeGroupTransform), err.Error())
}

func NewGroupToJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeGroupToJSON), err.Error())
}

func NewGroupDecodeJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeGroupDecodeJSON), err.Error())
}
func NewWalletTransformError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeWalletTransform), err.Error())
}

func NewWalletToJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeWalletToJSON), err.Error())
}

func NewWalletDecodeJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeWalletDecodeJSON), err.Error())
}
func NewAmountTransformError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeAmountTransform), err.Error())
}

func NewAmountToJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeAmountToJSON), err.Error())
}

func NewAmountDecodeJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeAmountDecodeJSON), err.Error())
}
