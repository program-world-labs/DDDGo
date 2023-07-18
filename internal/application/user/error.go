package user

import (
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeRepository    = domainerrors.ErrorCodeAdapterHTTPUser + iota // 20000000
	ErrorCodeValidateInput                                                // 20000001
	ErrorCodeCast                                                         // 20000002
)

var (
	ErrValidation         = errors.New("validation failed")
	ErrCastToEntityFailed = errors.New("cast to entity failed")
)
