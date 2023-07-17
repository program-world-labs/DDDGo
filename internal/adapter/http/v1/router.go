// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/program-world-labs/pwlogger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/program-world-labs/DDDGo/config"
	"github.com/program-world-labs/DDDGo/docs"
	"github.com/program-world-labs/DDDGo/internal/adapter/http/v1/role"
	"github.com/program-world-labs/DDDGo/internal/adapter/http/v1/user"
	application_role "github.com/program-world-labs/DDDGo/internal/application/role"
	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
)

type Services struct {
	User application_user.IService
	Role application_role.IService
}

// NewRouter -.
// Swagger spec:
// @title       AI Service API
// @description Using AI to do something.
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
// Swagger base path.
func NewRouter(l pwlogger.Interface, s Services, cfg *config.Config) *gin.Engine {
	handler := gin.New()
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(otelgin.Middleware(cfg.App.Name))

	// Swagger
	docs.SwaggerInfo.Version = cfg.App.Version
	docs.SwaggerInfo.Title = cfg.App.Name

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
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
