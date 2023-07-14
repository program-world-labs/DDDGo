package role

import (
	"github.com/rs/zerolog"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeRoleUsecase       = domainerrors.ErrorCodeAdapterRole + iota // 1000000
	ErrorCodeRoleBindJSON                                                 // 1000001
	ErrorCodeRoleCopyToInput                                              // 1000002
	ErrorCodeRoleBindQuery                                                // 1000003
	ErrorCodeRoleValidateInput                                            // 1000004
	ErrorCodeRoleBindUri                                                  // 1000005
)

type ErrorEvent struct {
	err error
}

func (a ErrorEvent) MarshalZerologObject(e *zerolog.Event) {
	e.Err(a.err).Str("event", "RoleError")
}
