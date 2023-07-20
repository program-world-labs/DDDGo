package role

import (
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
)

const (
	ErrorCodeRepository = domainerrors.ErrorCodeApplicationRole + iota
	ErrorCodeValidateInput
	ErrorCodeCast
	ErrorCodeApplyEvent
	ErrorCodePublishEvent
)

var (
	ErrValidation         = errors.New("validation failed")
	ErrCastToEntityFailed = errors.New("cast to entity failed")
)

// func NewRepositoryError(err error) *domainerrors.ErrorInfo {
// 	var repoError *domainerrors.ErrorInfo
// 	if errors.As(err, &repoError) {
// 		code, atoiErr := strconv.Atoi(repoError.Code)
// 		if atoiErr != nil {
// 			code = 0
// 		}

// 		return domainerrors.New(fmt.Sprint(ErrorCodeRepository+code), err.Error())
// 	}

// 	return domainerrors.New(fmt.Sprint(ErrorCodeRepository), err.Error())
// }

// func NewValidateInputError(err error) *domainerrors.ErrorInfo {
// 	errValidateInput := domainerrors.New(fmt.Sprint(ErrorCodeValidateInput), err.Error())

// 	return errValidateInput
// }

// func NewCastError(err error) *domainerrors.ErrorInfo {
// 	errCast := domainerrors.New(fmt.Sprint(ErrorCodeCast), err.Error())

// 	return errCast
// }
