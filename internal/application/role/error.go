package role

import (
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeRepository = domainerrors.ErrorCodeApplicationRole + domainerrors.ErrorCodeApplicationRole + iota
	ErrorCodeValidateInput
	ErrorCodeCast
	ErrorCodeApplyEvent
	ErrorCodePublishEvent
)

var (
	ErrValidation         = errors.New("validation failed")
	ErrCastToEntityFailed = errors.New("cast to entity failed")
)
