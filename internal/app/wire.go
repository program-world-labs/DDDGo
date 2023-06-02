//go:build wireinject
// +build wireinject

package app

import (
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/program-world-labs/DDDGo/config"
	v1 "github.com/program-world-labs/DDDGo/internal/adapter/http/v1"
	usecase "github.com/program-world-labs/DDDGo/internal/application/user"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource/cache"
	repo "github.com/program-world-labs/DDDGo/internal/infra/datasource/sql"
	"github.com/program-world-labs/DDDGo/internal/infra/repository"
	"github.com/program-world-labs/DDDGo/pkg/cache/local"
	redis_cache "github.com/program-world-labs/DDDGo/pkg/cache/redis"
	"github.com/program-world-labs/DDDGo/pkg/httpserver"
	"github.com/program-world-labs/DDDGo/pkg/logger"
	sqlgorm "github.com/program-world-labs/DDDGo/pkg/sql_gorm"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func provideLogger(cfg *config.Config) (logger.Interface, error) {
	return logger.New(cfg.Log.Level), nil
}

func providePostgres(cfg *config.Config) (*gorm.DB, error) {
	client, err := sqlgorm.New(cfg.PG.URL, sqlgorm.MaxPoolSize(cfg.PG.PoolMax))
	return client.DB, err
}

func provideRedisCache(cfg *config.Config) (*redis.Client, error) {
	cache, err := redis_cache.New(cfg.Redis.DSN)
	return cache.Client, err
}

func provideLocalCache() (*bigcache.BigCache, error) {
	cache, err := local.New()
	return cache.Client, err
}

func provideUserRepo(sqlDatasource *repo.UserDatasourceImpl, redisCacheDatasource *cache.RedisCacheDataSourceImpl, bigCacheDatasource *cache.BigCacheDataSourceImpl) *repository.UserRepoImpl {
	return repository.NewUserRepoImpl(sqlDatasource, redisCacheDatasource, bigCacheDatasource)
}

func provideUserUseCase(userRepo *repository.UserRepoImpl) usecase.IUserUseCase {
	return usecase.NewUserUseCaseImpl(userRepo)
}

func provideHTTPServer(handler *gin.Engine, cfg *config.Config) *httpserver.Server {
	return httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}

var appSet = wire.NewSet(
	provideLogger,
	providePostgres,
	provideRedisCache,
	provideLocalCache,
	repo.NewUserDatasourceImpl,
	cache.NewRedisCacheDataSourceImpl,
	cache.NewBigCacheDataSourceImp,
	provideUserRepo,
	provideUserUseCase,
	v1.NewRouter,
	provideHTTPServer,
)

func InitializeHTTPServer(cfg *config.Config) (*httpserver.Server, error) {
	wire.Build(appSet)
	return &httpserver.Server{}, nil
}
