package domainerrors

import (
	"encoding/json"

	"github.com/pkg/errors"
)

const (
	GruopID                          = "A"
	ErrorCodeAdapterRole             = 10000000
	ErrorCodeAdapterUser             = 20000000
	ErrorCodeApplicationRole         = 1000000
	ErrorCodeDomainRole              = 100000
	ErrorCodeRepo                    = 10000
	ErrorCodeDatasourceSQL           = 1000
	ErrorCodeDatasourceCache         = 2000
	ErrorCodeDatasourceUserRepoDTO   = 3000
	ErrorCodeDatasourceRoleRepoDTO   = 4000
	ErrorCodeDatasourceGroupRepoDTO  = 5000
	ErrorCodeDatasourceWalletRepoDTO = 6000
	ErrorCodeDatasourceAmountRepoDTO = 7000
	ErrorCodeSystem                  = "00000000"
)

type ErrorInfo struct {
	Code    string      `json:"code"`
	Message string      `json:"error"`
	Err     interface{} `json:"-"`
}

// ErrorInfo implements the error interface.
func (e *ErrorInfo) Error() string {
	return e.Message
}

func (e *ErrorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code    string `json:"code"`
		Message string `json:"error"`
	}{
		Code:    e.Code,
		Message: e.Message,
	})
}

// New returns a new error with an error code and error message.
func New(code string, msg string) *ErrorInfo {
	return &ErrorInfo{Code: code, Message: msg}
}

// Cause returns the underlying cause of an error, if available.
func Cause(err error) error {
	return errors.Cause(err)
}

// IsErrorCode returns true if the given error has the given error code.
func IsErrorCode(err error, code string) bool {
	var e *ErrorInfo
	if errors.As(err, &e) {
		return e.Code == code
	}

	return false
}
