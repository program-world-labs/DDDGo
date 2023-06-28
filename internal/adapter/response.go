package adapter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	domain_errors "github.com/program-world-labs/DDDGo/internal/domain/errors"
)

type Response struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func HandleErrorResponse(c *gin.Context, err error) {
	// Check if the error is of type domain_errors.ErrorInfo
	var errorInfo domain_errors.ErrorInfo
	if errors.As(err, &errorInfo) {
		c.JSON(http.StatusOK, errorInfo)

		return
	}

	// If the error is not of type domain_errors.ErrorInfo, create a new ErrorInfo struct
	errorInfo = domain_errors.ErrorInfo{
		Code:    domain_errors.ErrorCodeInternalServerError,
		Message: domain_errors.GetErrorMessage(domain_errors.ErrorCodeInternalServerError),
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
