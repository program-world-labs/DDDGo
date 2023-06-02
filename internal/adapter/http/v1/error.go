package v1

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error" example:"message"`
}

func (*ErrorResponse) Errors(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, ErrorResponse{msg})
}
