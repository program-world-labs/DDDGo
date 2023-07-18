package role

import (
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeRoleUsecase = domainerrors.ErrorCodeAdapterMessageRole + iota
	ErrorCodeRoleCopyToInput
	ErrorCodeHandleMessage
)
