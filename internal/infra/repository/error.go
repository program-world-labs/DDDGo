package repository

import (
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeDatasource = domainerrors.ErrorCodeRepo + iota
	ErrorCodeRepoTransform
	ErrorCodeRepoBackToDomain
	ErrorCodeRepoCast
	ErrorCodeRepoCreate
	ErrorCodeRepoDelete
	ErrorCodeRepoUpdate
	ErrorCodeRepoUpdateWithFields
	ErrorCodeRepoGet
	ErrorCodeRepoGetAll
	ErrorCodeRepoSet
	ErrorCodeRepoCreateTx
	ErrorCodeRepoDeleteTx
	ErrorCodeRepoUpdateTx
	ErrorCodeRepoUpdateWithFieldsTx
	ErrorCodeRepoParseMap
)

// func NewDatasourceError(err error) *domainerrors.ErrorInfo {
// 	var repoError *domainerrors.ErrorInfo
// 	if errors.As(err, &repoError) {
// 		code, atoiErr := strconv.Atoi(repoError.Code)
// 		if atoiErr != nil {
// 			code = 0
// 		}

// 		return domainerrors.New(fmt.Sprint(ErrorCodeDatasource+code), err.Error())
// 	}

// 	return domainerrors.New(fmt.Sprint(ErrorCodeDatasource), err.Error())
// }
