package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
	common_error "github.com/program-world-labs/DDDGo/internal/domain/errors"
)

type Response struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

// func ErrorResponse(c *gin.Context, code common_error.ErrorCode) {
// 	c.JSON(http.StatusOK, common_error.GetErrorInfo(code))
// }

func HandleErrorResponse(c *gin.Context, err error) {
	// Convert the error to an ErrorInfo struct for JSON serialization
	errorInfo, ok := err.(common_error.ErrorInfo)
	if !ok {
		errorInfo = common_error.ErrorInfo{
			Code:    common_error.ErrorCodeInternalServerError,
			Message: common_error.GetErrorMessage(common_error.ErrorCodeInternalServerError)}
	}
	c.JSON(http.StatusOK, errorInfo)
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:  0,
		Error: "",
		Data:  data,
	})
}
