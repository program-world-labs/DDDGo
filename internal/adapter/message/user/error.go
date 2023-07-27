package user

import (
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeAdapterMessageUser = domainerrors.ErrorCodeAdapter + domainerrors.ErrorCodeAdapterMessage + domainerrors.ErrorCodeAdapterUser + iota
	ErrorCodeCopyToInput
	ErrorCodeHandleMessage
)
