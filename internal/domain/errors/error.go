package errors

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type ErrorCode int

const (
	// Controller Errors
	// 參數錯誤.
	ErrorCodeInvalidParameter ErrorCode = 1000 + iota
	// 參數驗證錯誤.
	ErrorCodeParamValidationFailed
	// 認證錯誤 - Authorization header is missing.
	ErrorCodeAuthHeaderMissing
	// Claim UID is missing.
	ErrorCodeClaimUIDMissing
	// Claim HospitalID is missing.
	ErrorCodeClaimHospitalIDMissing
	// SearchQuery is missing.
	ErrorCodeSearchQueryMissing
	// 參數遺失.
	ErrorCodeParamMissing
	// 參數格式錯誤.
	ErrorCodeParamFormatInvalid
	// 資源不存在.
	ErrorCodeResourceNotFound
	// 權限不足.
	ErrorCodePermissionDenied
	// Authorization header is invalid.
	ErrorCodeAuthHeaderInvalid
	// Role type is invalid.
	ErrorCodeRoleTypeInvalid
	// 運算錯誤.
	ErrorCodeDivisionByZero
	// 資料欄位轉換錯誤.
	ErrorCodeDataFieldConvertError

	// Entity Errors.
	ErrorCodeEntityParamInvalid ErrorCode = 2000 + iota
	ErrorCodeEntityValidationFailed
	ErrorCodeEntityFieldHasExist
	ErrorCodeEntityFieldNotFound

	// UseCase Errors.
	ErrorCodeUseCaseParamInvalid ErrorCode = 3000 + iota
	ErrorCodeUseCaseValidationFailed
	ErrorCodeUseCaseDataNotFound
	ErrorCodeUseCaseTransactionFailed
	ErrorCodeUseCaseOutputFailed
	ErrorCodeUseCasePermissionDenied
	ErrorCodeUseCasePublishFailed
	ErrorCodeUseCaseProcessFailed
	ErrorCodeUseCaseTagIsBinding

	// Database Errors.
	ErrorCodeDataSourceDeleteFailed ErrorCode = 4000 + iota
	ErrorCodeDataSourceCreateFailed
	ErrorCodeDataSourceUpdateFailed
	ErrorCodeDataSourceQueryFailed
	ErrorCodeDataSourceConvertFieldFailed
	ErrorCodeDataSourceUploadFailed
	ErrorCodeDataSourceDownloadFailed
	ErrorCodeDataSourceDataNotFound

	// Repository Errors.
	ErrorCodeRepositoryCreateFailed ErrorCode = 5000 + iota
	ErrorCodeRepositoryUpdateFailed
	ErrorCodeRepositoryDeleteFailed
	ErrorCodeRepositoryQueryFailed

	// Pub/Sub Errors.
	ErrorCodePubSubPublishFailed ErrorCode = 6000 + iota
	ErrorCodePubSubSubscribeFailed
	ErrorCodePubSubReceivedFailed
	ErrorCodePubSubCreateTopicFailed
	ErrorCodePubSubRouteFailed
	ErrorCodePubSubHandlerFailed
	ErrorCodePubSubDataFormatInvalid
	ErrorCodePubSubShutdownFailed

	// Internal Errors
	// 內部錯誤.
	ErrorCodeInternalServerError ErrorCode = 5000 + iota
)

type ErrorInfo struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"error"`
	Err     interface{} `json:"-"`
}

//nolint:gochecknoglobals // TODO. 暫時先這樣，後續再改.
var errorInfoMap = map[ErrorCode]ErrorInfo{
	// Controller Errors
	ErrorCodeInvalidParameter:       {Code: ErrorCodeInvalidParameter, Message: "Invalid parameter"},
	ErrorCodeParamValidationFailed:  {Code: ErrorCodeParamValidationFailed, Message: "Param validation failed"},
	ErrorCodeAuthHeaderMissing:      {Code: ErrorCodeAuthHeaderMissing, Message: "Authorization header is missing"},
	ErrorCodeClaimUIDMissing:        {Code: ErrorCodeClaimUIDMissing, Message: "Claim UID is missing"},
	ErrorCodeClaimHospitalIDMissing: {Code: ErrorCodeClaimHospitalIDMissing, Message: "Claim HospitalID is missing"},
	ErrorCodeParamMissing:           {Code: ErrorCodeParamMissing, Message: "Param is missing"},
	ErrorCodeParamFormatInvalid:     {Code: ErrorCodeParamFormatInvalid, Message: "Param format is invalid"},
	ErrorCodeSearchQueryMissing:     {Code: ErrorCodeSearchQueryMissing, Message: "SearchQuery is missing"},
	ErrorCodeResourceNotFound:       {Code: ErrorCodeResourceNotFound, Message: "Resource not found"},
	ErrorCodePermissionDenied:       {Code: ErrorCodePermissionDenied, Message: "Permission denied"},
	ErrorCodeAuthHeaderInvalid:      {Code: ErrorCodeAuthHeaderInvalid, Message: "Authorization header is invalid"},
	ErrorCodeRoleTypeInvalid:        {Code: ErrorCodeRoleTypeInvalid, Message: "Role permission is invalid"},
	ErrorCodeDivisionByZero:         {Code: ErrorCodeDivisionByZero, Message: "Division by zero"},
	ErrorCodeDataFieldConvertError:  {Code: ErrorCodeDataFieldConvertError, Message: "Data field convert error"},
	// Enttity Errors
	ErrorCodeEntityParamInvalid:     {Code: ErrorCodeEntityParamInvalid, Message: "Entity param invalid"},
	ErrorCodeEntityValidationFailed: {Code: ErrorCodeEntityValidationFailed, Message: "Entity validation failed"},
	ErrorCodeEntityFieldHasExist:    {Code: ErrorCodeEntityFieldHasExist, Message: "Entity field has exist"},
	ErrorCodeEntityFieldNotFound:    {Code: ErrorCodeEntityFieldNotFound, Message: "Entity field not found"},
	// UseCase Errors
	ErrorCodeUseCaseParamInvalid:      {Code: ErrorCodeUseCaseParamInvalid, Message: "UseCase param invalid"},
	ErrorCodeUseCaseValidationFailed:  {Code: ErrorCodeUseCaseValidationFailed, Message: "UseCase validation failed"},
	ErrorCodeUseCaseDataNotFound:      {Code: ErrorCodeUseCaseDataNotFound, Message: "UseCase data not found"},
	ErrorCodeUseCaseTransactionFailed: {Code: ErrorCodeUseCaseTransactionFailed, Message: "UseCase transaction failed"},
	ErrorCodeUseCaseOutputFailed:      {Code: ErrorCodeUseCaseOutputFailed, Message: "UseCase output failed"},
	ErrorCodeUseCasePermissionDenied:  {Code: ErrorCodeUseCasePermissionDenied, Message: "UseCase permission denied"},
	ErrorCodeUseCasePublishFailed:     {Code: ErrorCodeUseCasePublishFailed, Message: "UseCase publish failed"},
	ErrorCodeUseCaseProcessFailed:     {Code: ErrorCodeUseCaseProcessFailed, Message: "UseCase process failed"},
	ErrorCodeUseCaseTagIsBinding:      {Code: ErrorCodeUseCaseTagIsBinding, Message: "UseCase tag is binding"},
	// Data Source Errors
	ErrorCodeDataSourceDeleteFailed:       {Code: ErrorCodeDataSourceDeleteFailed, Message: "Data source delete failed"},
	ErrorCodeDataSourceCreateFailed:       {Code: ErrorCodeDataSourceCreateFailed, Message: "Data source create failed"},
	ErrorCodeDataSourceUpdateFailed:       {Code: ErrorCodeDataSourceUpdateFailed, Message: "Data source update failed"},
	ErrorCodeDataSourceQueryFailed:        {Code: ErrorCodeDataSourceQueryFailed, Message: "Data source query failed"},
	ErrorCodeDataSourceConvertFieldFailed: {Code: ErrorCodeDataSourceConvertFieldFailed, Message: "Data source convert field failed"},
	ErrorCodeDataSourceUploadFailed:       {Code: ErrorCodeDataSourceUploadFailed, Message: "Data source upload failed"},
	ErrorCodeDataSourceDownloadFailed:     {Code: ErrorCodeDataSourceDownloadFailed, Message: "Data source download failed"},
	ErrorCodeDataSourceDataNotFound:       {Code: ErrorCodeDataSourceDataNotFound, Message: "Data source data not found"},

	// Repository Errors
	ErrorCodeRepositoryCreateFailed: {Code: ErrorCodeRepositoryCreateFailed, Message: "Repository create failed"},
	ErrorCodeRepositoryUpdateFailed: {Code: ErrorCodeRepositoryUpdateFailed, Message: "Repository update failed"},
	ErrorCodeRepositoryDeleteFailed: {Code: ErrorCodeRepositoryDeleteFailed, Message: "Repository delete failed"},
	ErrorCodeRepositoryQueryFailed:  {Code: ErrorCodeRepositoryQueryFailed, Message: "Repository query failed"},

	// Pub/Sub Errors
	ErrorCodePubSubPublishFailed:     {Code: ErrorCodePubSubPublishFailed, Message: "Pub/Sub publish failed"},
	ErrorCodePubSubSubscribeFailed:   {Code: ErrorCodePubSubSubscribeFailed, Message: "Pub/Sub subscribe failed"},
	ErrorCodePubSubReceivedFailed:    {Code: ErrorCodePubSubReceivedFailed, Message: "Pub/Sub received failed"},
	ErrorCodePubSubCreateTopicFailed: {Code: ErrorCodePubSubCreateTopicFailed, Message: "Pub/Sub create topic failed"},
	ErrorCodePubSubRouteFailed:       {Code: ErrorCodePubSubRouteFailed, Message: "Pub/Sub route failed"},
	ErrorCodePubSubHandlerFailed:     {Code: ErrorCodePubSubHandlerFailed, Message: "Pub/Sub handler failed"},
	ErrorCodePubSubDataFormatInvalid: {Code: ErrorCodePubSubDataFormatInvalid, Message: "Pub/Sub data format invalid"},
	ErrorCodePubSubShutdownFailed:    {Code: ErrorCodePubSubShutdownFailed, Message: "Pub/Sub shutdown failed"},

	// Infrastructure Errors
	// Internal Errors
	ErrorCodeInternalServerError: {Code: ErrorCodeInternalServerError, Message: "Internal server error"},
}

// ErrorInfo implements the error interface.
func (e ErrorInfo) Error() string {
	return e.Message
}

func (e *ErrorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code    ErrorCode `json:"code"`
		Message string    `json:"error"`
	}{
		Code:    e.Code,
		Message: e.Message,
	})
}

// New returns a new error with an error code and error message.
func New(code ErrorCode) ErrorInfo {
	return ErrorInfo{Code: code, Message: errorInfoMap[code].Message}
}

// Wrap returns a new error with an error code and error message, wrapping an existing error.
func Wrap(code ErrorCode, err error) *ErrorInfo {
	return &ErrorInfo{Code: code, Message: GetErrorMessage(code), Err: errors.WithStack(err)}
}

// Cause returns the underlying cause of an error, if available.
func Cause(err error) error {
	return errors.Cause(err)
}

// IsErrorCode returns true if the given error has the given error code.
func IsErrorCode(err error, code ErrorCode) bool {
	var e *ErrorInfo
	if errors.As(err, &e) {
		return e.Code == code
	}

	return false
}

func GetErrorInfo(code ErrorCode) ErrorInfo {
	errorInfo, ok := errorInfoMap[code]
	if !ok {
		return errorInfoMap[ErrorCodeInternalServerError]
	}

	return errorInfo
}

func GetErrorMessage(code ErrorCode) string {
	return GetErrorInfo(code).Message
}
