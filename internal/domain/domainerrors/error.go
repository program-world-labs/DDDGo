package domainerrors

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	GruopID = "PAAC"
	// Adapter Layer Error Code Offset.
	ErrorCodeAdapter         = 100000000
	ErrorCodeAdapterHTTP     = 010000000
	ErrorCodeAdapterMessage  = 020000000
	ErrorCodeAdapterUser     = 001000000
	ErrorCodeAdapterRole     = 002000000
	ErrorCodeAdapterGroup    = 003000000
	ErrorCodeAdapterWallet   = 004000000
	ErrorCodeAdapterCurrency = 005000000
	// Application Layer Error Code Offset.
	ErrorCodeApplication         = 200000000
	ErrorCodeApplicationUser     = 010000000
	ErrorCodeApplicationRole     = 020000000
	ErrorCodeApplicationGroup    = 030000000
	ErrorCodeApplicationWallet   = 040000000
	ErrorCodeApplicationCurrency = 050000000
	// Domain Layer Error Code Offset.
	ErrorCodeDomain         = 300000000
	ErrorCodeDomainEntity   = 010000000
	ErrorCodeDomainEvent    = 020000000
	ErrorCodeDomainUser     = 001000000
	ErrorCodeDomainRole     = 002000000
	ErrorCodeDomainGroup    = 003000000
	ErrorCodeDomainWallet   = 004000000
	ErrorCodeDomainCurrency = 005000000
	// Infra Layer Error Code Offset.
	ErrorCodeInfra           = 200000000
	ErrorCodeInfraDatasource = 010000000
	ErrorCodeInfraDTO        = 020000000
	ErrorCodeInfraRepo       = 030000000
	// Infra Layer Datasource Error Code.
	ErrorCodeInfraDatasourceSQL        = 001000000
	ErrorCodeInfraDatasourceCache      = 002000000
	ErrorCodeInfraDatasourceEventStore = 003000000
	// Infra Layer DTO Error Code.
	ErrorCodeInfraDTOMapper = 001000000
	ErrorCodeInfraDTOVO     = 002000000
	ErrorCodeInfraDTOBase   = 003000000
	ErrorCodeInfraDTOList   = 004000000
	// Infra Layer Repository Error Code.
	ErrorCodeInfraRepoCRUD        = 001000000
	ErrorCodeInfraRepoTransaction = 002000000
	ErrorCodeInfraRepoUser        = 003000000
	ErrorCodeInfraRepoRole        = 004000000
	ErrorCodeInfraRepoGroup       = 005000000
	ErrorCodeInfraRepoWallet      = 006000000
	ErrorCodeInfraRepoCurrency    = 007000000

	ErrorCodeSystem = 900000000
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
	if errorCode >= ErrorCodeSystem {
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
