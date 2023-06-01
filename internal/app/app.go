// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/program-world-labs/DDDGo/config"
	v1 "github.com/program-world-labs/DDDGo/internal/adapter/http/v1"
	usecase "github.com/program-world-labs/DDDGo/internal/application/user"
	repo "github.com/program-world-labs/DDDGo/internal/infra/datasource/sql"
	"github.com/program-world-labs/DDDGo/internal/infra/repository"
	"github.com/program-world-labs/DDDGo/pkg/httpserver"
	"github.com/program-world-labs/DDDGo/pkg/logger"
	sqlgorm "github.com/program-world-labs/DDDGo/pkg/sql_gorm"
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
