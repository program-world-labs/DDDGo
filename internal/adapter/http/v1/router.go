// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/program-world-labs/pwlogger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Blank import removed as it is not justified
	// _ "github.com/program-world-labs/DDDGo/docs".
	_ "github.com/program-world-labs/DDDGo/docs"
	"github.com/program-world-labs/DDDGo/internal/adapter/http/v1/role"
	"github.com/program-world-labs/DDDGo/internal/adapter/http/v1/user"
	"github.com/program-world-labs/DDDGo/internal/application"
)

// type Services struct {
// 	User application_user.IService
// 	Role application_role.IService
// }

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
// Swagger base path.
func NewRouter(l pwlogger.Interface, s application.Services) *gin.Engine {
	handler := gin.New()
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		user.NewUserRoutes(h, s.User, l)
		role.NewRoleRoutes(h, s.Role, l)
	}

	return handler
}
