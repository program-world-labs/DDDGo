package eventstore

import (
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeCacheSet = domainerrors.ErrorCodeInfraDatasource + domainerrors.ErrorCodeInfraDatasource + domainerrors.ErrorCodeInfraDatasourceEventStore + iota
	ErrorCodeResourceNotFound
	ErrorCodeEventFormatWrong
)
