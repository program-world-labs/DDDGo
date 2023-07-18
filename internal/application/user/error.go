package user

import (
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeRepository    = domainerrors.ErrorCodeApplicationUser + iota // 100000
	ErrorCodeValidateInput                                                // 100001
	ErrorCodeCast                                                         // 100002
)

var (
	ErrValidation         = errors.New("validation failed")
	ErrCastToEntityFailed = errors.New("cast to entity failed")
)
