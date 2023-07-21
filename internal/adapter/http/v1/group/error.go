package group

import (
	"github.com/rs/zerolog"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeGroupUsecase = domainerrors.ErrorCodeAdapterHTTPGroup + iota
	ErrorCodeGroupBindJSON
	ErrorCodeGroupCopyToInput
	ErrorCodeGroupBindQuery
	ErrorCodeGroupValidateInput
	ErrorCodeGroupBindURI
)

type ErrorEvent struct {
	err error
}

func (a ErrorEvent) MarshalZerologObject(e *zerolog.Event) {
	e.Err(a.err).Str("event", "GroupError")
}
