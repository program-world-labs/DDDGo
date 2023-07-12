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
	ErrorCodeUserInvalidFilterField
	ErrorCodeUserInvalidOrderField
	ErrorCodeUserParseMap
)
const (
	ErrorCodeRoleTransform = domainerrors.ErrorCodeDatasourceRoleRepoDTO + iota + 1
	ErrorCodeRoleBackToDomain
	ErrorCodeRoleToJSON
	ErrorCodeRoleDecodeJSON
	ErrorCodeRoleInvalidFilterField
	ErrorCodeRoleInvalidOrderField
	ErrorCodeRoleParseMap
)
const (
	ErrorCodeGroupTransform = domainerrors.ErrorCodeDatasourceGroupRepoDTO + iota + 1
	ErrorCodeGroupBackToDomain
	ErrorCodeGroupToJSON
	ErrorCodeGroupDecodeJSON
	ErrorCodeGroupInvalidFilterField
	ErrorCodeGroupInvalidOrderField
	ErrorCodeGroupParseMap
)
const (
	ErrorCodeWalletTransform = domainerrors.ErrorCodeDatasourceWalletRepoDTO + iota + 1
	ErrorCodeWalletBackToDomain
	ErrorCodeWalletToJSON
	ErrorCodeWalletDecodeJSON
	ErrorCodeWalletInvalidFilterField
	ErrorCodeWalletInvalidOrderField
	ErrorCodeWalletParseMap
)
const (
	ErrorCodeAmountTransform = domainerrors.ErrorCodeDatasourceAmountRepoDTO + iota + 1
	ErrorCodeAmountBackToDomain
	ErrorCodeAmountToJSON
	ErrorCodeAmountDecodeJSON
	ErrorCodeAmountInvalidFilterField
	ErrorCodeAmountInvalidOrderField
	ErrorCodeAmountParseMap
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

func NewUserInvalidError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeUserInvalidFilterField), err.Error())
}

func NewUserParseMapError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeUserParseMap), err.Error())
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

func NewRoleInvalidError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeRoleInvalidFilterField), err.Error())
}

func NewRoleParseMapError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeRoleParseMap), err.Error())
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

func NewGroupInvalidError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeGroupInvalidFilterField), err.Error())
}

func NewGroupParseMapError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeGroupParseMap), err.Error())
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

func NewWalletInvalidError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeWalletInvalidFilterField), err.Error())
}

func NewWalletParseMapError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeWalletParseMap), err.Error())
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

func NewAmountInvalidError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeAmountInvalidFilterField), err.Error())
}

func NewAmountParseMapError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeAmountParseMap), err.Error())
}
