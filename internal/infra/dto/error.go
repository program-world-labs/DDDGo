package dto

import (
	"errors"

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
	ErrorCodeWalletBalanceTransform = domainerrors.ErrorCodeDatasourceWalletBalanceRepoDTO + iota + 1
	ErrorCodeWalletBalanceBackToDomain
	ErrorCodeWalletBalanceToJSON
	ErrorCodeWalletBalanceDecodeJSON
	ErrorCodeWalletBalanceInvalidFilterField
	ErrorCodeWalletBalanceInvalidOrderField
	ErrorCodeWalletBalanceParseMap
)
const (
	ErrorCodeCurrencyTransform = domainerrors.ErrorCodeDatasourceCurrencyRepoDTO + iota + 1
	ErrorCodeCurrencyBackToDomain
	ErrorCodeCurrencyToJSON
	ErrorCodeCurrencyDecodeJSON
	ErrorCodeCurrencyInvalidFilterField
	ErrorCodeCurrencyInvalidOrderField
	ErrorCodeCurrencyParseMap
)

var (
	ErrParesMapFailed = errors.New("parse map failed")
)
