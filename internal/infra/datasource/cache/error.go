package cache

import (
	"fmt"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeCacheSet = domainerrors.ErrorCodeDatasourceCache + iota + 1
	ErrorCodeCacheDelete
	ErrorCodeCacheGet
)

func NewSetError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeCacheSet), err.Error())
}

func NewDeleteError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeCacheDelete), err.Error())
}

func NewGetError(err error) *domainerrors.ErrorInfo {
	return domainerrors.New(fmt.Sprint(ErrorCodeCacheGet), err.Error())
}
