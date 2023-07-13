package cache

import (
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeCacheSet = domainerrors.ErrorCodeDatasourceCache + iota + 1
	ErrorCodeCacheDelete
	ErrorCodeCacheGet
)
