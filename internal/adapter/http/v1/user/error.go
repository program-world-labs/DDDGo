package user

import (
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/rs/zerolog"
)

const (
	ErrorCodeUserUsecase       = domainerrors.ErrorCodeAdapterUser + iota // 2000000
	ErrorCodeUserBindJSON                                                 // 2000001
	ErrorCodeUserCopyToInput                                              // 2000002
	ErrorCodeUserBindQuery                                                // 2000003
	ErrorCodeUserValidateInput                                            // 2000004
	ErrorCodeUserBindURI                                                  // 2000005
)

type ErrorEvent struct {
	err error
}

func (a ErrorEvent) MarshalZerologObject(e *zerolog.Event) {
	e.Err(a.err).Str("event", "UserError")
}
