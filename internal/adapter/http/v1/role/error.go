package role

import (
	"errors"
	"fmt"
	"strconv"

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

func NewBindJSONError(err error) *domainerrors.ErrorInfo {
	errBindJSON := domainerrors.New(domainerrors.GruopID+fmt.Sprint(ErrorCodeRoleBindJSON), err.Error())

	return errBindJSON
}

func NewBindQueryError(err error) *domainerrors.ErrorInfo {
	errBindQuery := domainerrors.New(domainerrors.GruopID+fmt.Sprint(ErrorCodeRoleBindQuery), err.Error())

	return errBindQuery
}

func NewBindUriError(err error) *domainerrors.ErrorInfo {
	errBindUri := domainerrors.New(domainerrors.GruopID+fmt.Sprint(ErrorCodeRoleBindUri), err.Error())

	return errBindUri
}

func NewValidateInputError(err error) *domainerrors.ErrorInfo {
	errValidateInput := domainerrors.New(domainerrors.GruopID+fmt.Sprint(ErrorCodeRoleValidateInput), err.Error())

	return errValidateInput
}

func NewCopyError(err error) *domainerrors.ErrorInfo {
	errCopyToInput := domainerrors.New(domainerrors.GruopID+fmt.Sprint(ErrorCodeRoleCopyToInput), err.Error())

	return errCopyToInput
}

func NewUsecaseError(err error) *domainerrors.ErrorInfo {
	var usecaseError *domainerrors.ErrorInfo
	if errors.As(err, &usecaseError) {
		code, atoiErr := strconv.Atoi(usecaseError.Code)
		if atoiErr != nil {
			code = 0
		}

		return domainerrors.New(domainerrors.GruopID+fmt.Sprint(ErrorCodeRoleUsecase+code), err.Error())
	}

	return domainerrors.New(domainerrors.GruopID+fmt.Sprint(ErrorCodeRoleUsecase), err.Error())
}

type ErrorEvent struct {
	err error
}

func (a ErrorEvent) MarshalZerologObject(e *zerolog.Event) {
	e.Err(a.err).Str("event", "RoleError")
}
