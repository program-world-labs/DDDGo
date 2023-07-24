package role

import (
	"github.com/rs/zerolog"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeAdapterHTTPRole = domainerrors.ErrorCodeAdapter + domainerrors.ErrorCodeAdapterHTTP + domainerrors.ErrorCodeAdapterRole + iota
	ErrorCodeExecuteUsecase
	ErrorCodeBindJSON
	ErrorCodeCopyToInput
	ErrorCodeBindQuery
	ErrorCodeValidateInput
	ErrorCodeBindURI
)

type ErrorEvent struct {
	err error
}

func (a ErrorEvent) MarshalZerologObject(e *zerolog.Event) {
	e.Err(a.err).Str("event", "RoleError")
}
