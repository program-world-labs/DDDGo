package domainerrors

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	GruopID                          = "LLM"
	ErrorCodeAdapterRole             = 10000000
	ErrorCodeAdapterUser             = 20000000
	ErrorCodeApplicationRole         = 1000000
	ErrorCodeApplicationUser         = 2000000
	ErrorCodeDomainRole              = 100000
	ErrorCodeRepo                    = 10000
	ErrorCodeDatasourceSQL           = 1000
	ErrorCodeDatasourceCache         = 2000
	ErrorCodeDatasourceUserRepoDTO   = 3000
	ErrorCodeDatasourceRoleRepoDTO   = 4000
	ErrorCodeDatasourceGroupRepoDTO  = 5000
	ErrorCodeDatasourceWalletRepoDTO = 6000
	ErrorCodeDatasourceAmountRepoDTO = 7000
	ErrorCodeSystem                  = 90000000
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

// Wrap returns a new error with an error code and error message, wrapping an existing error.
func Wrap(errorCode int, err error) *ErrorInfo {
	// Check if error code is adapter error code
	var group = ""
	if errorCode >= ErrorCodeAdapterRole {
		group = GruopID
	}

	// Check if err is already a ErrorInfo
	var e *ErrorInfo
	if errors.As(err, &e) {
		code, atoiErr := strconv.Atoi(e.Code)
		if atoiErr != nil {
			code = 0
		}

		return &ErrorInfo{
			Code:    group + strconv.Itoa(errorCode+code),
			Message: e.Message,
			Err:     errors.WithStack(err)}
	}

	return &ErrorInfo{Code: group + strconv.Itoa(errorCode), Message: err.Error(), Err: errors.WithStack(err)}
}

func WrapWithSpan(errorCode int, err error, span trace.Span) *ErrorInfo {
	span.SetStatus(codes.Error, err.Error())

	return Wrap(errorCode, err)
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
