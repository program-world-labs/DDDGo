package dto

import (
	"fmt"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeUserTransform = domainerrors.ErrorCodeDatasourceUserRepoDTO + iota + 1
	ErrorCodeUserBackToDomain
	ErrorCodeUserToJSON
	ErrorCodeUserDecodeJSON
)
const (
	ErrorCodeRoleTransform = domainerrors.ErrorCodeDatasourceRoleRepoDTO + iota + 1
	ErrorCodeRoleBackToDomain
	ErrorCodeRoleToJSON
	ErrorCodeRoleDecodeJSON
)
const (
	ErrorCodeGroupTransform = domainerrors.ErrorCodeDatasourceGroupRepoDTO + iota + 1
	ErrorCodeGroupBackToDomain
	ErrorCodeGroupToJSON
	ErrorCodeGroupDecodeJSON
)
const (
	ErrorCodeWalletTransform = domainerrors.ErrorCodeDatasourceWalletRepoDTO + iota + 1
	ErrorCodeWalletBackToDomain
	ErrorCodeWalletToJSON
	ErrorCodeWalletDecodeJSON
)
const (
	ErrorCodeAmountTransform = domainerrors.ErrorCodeDatasourceAmountRepoDTO + iota + 1
	ErrorCodeAmountBackToDomain
	ErrorCodeAmountToJSON
	ErrorCodeAmountDecodeJSON
)

func NewUserTransformError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeUserTransform), err.Error())
}

func NewUserBackToDomainError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeUserBackToDomain), err.Error())
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
func NewRoleBackToDomainError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeRoleBackToDomain), err.Error())
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
func NewGroupBackToDomainError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeGroupBackToDomain), err.Error())
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

func NewWalletBackToDomainError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeWalletBackToDomain), err.Error())
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

func NewAmountBackToDomainError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeAmountBackToDomain), err.Error())
}

func NewAmountToJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeAmountToJSON), err.Error())
}

func NewAmountDecodeJSONError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeAmountDecodeJSON), err.Error())
}
