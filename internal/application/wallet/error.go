package wallet

import (
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeRepository = domainerrors.ErrorCodeApplication + domainerrors.ErrorCodeApplicationWallet + iota
	ErrorCodeValidateInput
	ErrorCodeCast
)

var (
	ErrValidation         = errors.New("validation failed")
	ErrCastToEntityFailed = errors.New("cast to entity failed")
)
