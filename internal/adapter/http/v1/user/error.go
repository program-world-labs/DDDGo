package user

import (
	"github.com/rs/zerolog"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeAdapterHTTPUser = domainerrors.ErrorCodeAdapter + domainerrors.ErrorCodeAdapterHTTP + domainerrors.ErrorCodeAdapterUser + iota
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
	e.Err(a.err).Str("event", "UserError")
}
