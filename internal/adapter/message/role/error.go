package role

import (
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeAdapterMessageRole = domainerrors.ErrorCodeAdapter + domainerrors.ErrorCodeAdapterMessage + domainerrors.ErrorCodeAdapterRole + iota
	ErrorCodeCopyToInput
	ErrorCodeHandleMessage
)
