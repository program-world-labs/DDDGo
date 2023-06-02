package dto

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	common_error "github.com/program-world-labs/DDDGo/internal/domain/errors"
)

type Response struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func HandleErrorResponse(c *gin.Context, err error) {
	// Check if the error is of type common_error.ErrorInfo
	var errorInfo common_error.ErrorInfo
	if errors.As(err, &errorInfo) {
		c.JSON(http.StatusOK, errorInfo)

		return
	}

	// If the error is not of type common_error.ErrorInfo, create a new ErrorInfo struct
	errorInfo = common_error.ErrorInfo{
		Code:    common_error.ErrorCodeInternalServerError,
		Message: common_error.GetErrorMessage(common_error.ErrorCodeInternalServerError),
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
