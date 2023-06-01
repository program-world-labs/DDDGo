// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"gitlab.com/demojira/template.git/config"
	v1 "gitlab.com/demojira/template.git/internal/adapter/http/v1"
	usecase "gitlab.com/demojira/template.git/internal/application/user"
	repo "gitlab.com/demojira/template.git/internal/infra/datasource/sql"
	"gitlab.com/demojira/template.git/internal/infra/repository"
	"gitlab.com/demojira/template.git/pkg/httpserver"
	"gitlab.com/demojira/template.git/pkg/logger"
	sqlgorm "gitlab.com/demojira/template.git/pkg/sql_gorm"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Pkg - Postgres SQL
	pg, err := sqlgorm.New(cfg.PG.URL, sqlgorm.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Data source
	sqlDatasourceImpl := repo.NewCRUDDatasourceImpl(pg.DB)

	// User Repository
	userRepo := repository.NewUserRepoImpl(sqlDatasourceImpl, nil, nil)
	// Use case
	userUseCase := usecase.NewUserUseCaseImpl(
		userRepo,
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, userUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
