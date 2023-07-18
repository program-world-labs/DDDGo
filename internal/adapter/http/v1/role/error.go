package role

import (
	"github.com/rs/zerolog"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeRoleUsecase = domainerrors.ErrorCodeAdapterHTTPRole + iota
	ErrorCodeRoleBindJSON
	ErrorCodeRoleCopyToInput
	ErrorCodeRoleBindQuery
	ErrorCodeRoleValidateInput
	ErrorCodeRoleBindURI
)

type ErrorEvent struct {
	err error
}

func (a ErrorEvent) MarshalZerologObject(e *zerolog.Event) {
	e.Err(a.err).Str("event", "RoleError")
}
