package cache

import (
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeCacheSet = domainerrors.ErrorCodeInfraDatasource + domainerrors.ErrorCodeInfraDatasource + domainerrors.ErrorCodeInfraDatasourceCache + iota
	ErrorCodeCacheDelete
	ErrorCodeCacheGet
)
