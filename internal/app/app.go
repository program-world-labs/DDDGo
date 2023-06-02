// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/program-world-labs/DDDGo/config"
	"github.com/program-world-labs/DDDGo/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// // Pkg - Postgres SQL
	// pg, err := sqlgorm.New(cfg.PG.URL, sqlgorm.MaxPoolSize(cfg.PG.PoolMax))
	// if err != nil {
	// 	l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	// }
	// defer pg.Close()

	// redisCache, err := redis_cache.New(cfg.Redis.DSN)
	// if err != nil {
	// 	l.Fatal(fmt.Errorf("app - Run - redis.New: %w", err))
	// }
	// defer redisCache.Close()

	// bigCache, err := local.New()
	// if err != nil {
	// 	l.Fatal(fmt.Errorf("app - Run - bigcache.New: %w", err))
	// }
	// defer bigCache.Close()

	// // Data source
	// sqlDatasourceImpl := repo.NewUserDatasourceImpl(pg.DB)
	// redisCacheDatasourceImpl := cache.NewRedisCacheDataSourceImpl(redisCache.Client)
	// bigCacheDatasourceImpl := cache.NewBigCacheDataSourceImp(bigCache.Client)

	// // User Repository
	// userRepo := repository.NewUserRepoImpl(sqlDatasourceImpl, redisCacheDatasourceImpl, bigCacheDatasourceImpl)
	// // Use case
	// userUseCase := usecase.NewUserUseCaseImpl(
	// 	userRepo,
	// )

	// // HTTP Server
	// handler := v1.NewRouter(l, userUseCase)
	// httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	httpServer, err := InitializeHTTPServer(cfg)
	if err != nil {
		l.Error("failed to initialize HTTP server: %v", err)
	}

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
