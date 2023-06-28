//go:build wireinject
// +build wireinject

package app

import (
	"github.com/allegro/bigcache/v3"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/program-world-labs/pwlogger"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/config"
	v1 "github.com/program-world-labs/DDDGo/internal/adapter/http/v1"
	application_role "github.com/program-world-labs/DDDGo/internal/application/role"
	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource/cache"
	"github.com/program-world-labs/DDDGo/internal/infra/role"
	"github.com/program-world-labs/DDDGo/internal/infra/user"
	"github.com/program-world-labs/DDDGo/pkg/cache/local"
	redisCache "github.com/program-world-labs/DDDGo/pkg/cache/redis"
	"github.com/program-world-labs/DDDGo/pkg/httpserver"
	sqlgorm "github.com/program-world-labs/DDDGo/pkg/sql_gorm"
)

func providePostgres(cfg *config.Config) (*gorm.DB, error) {
	client, err := sqlgorm.New(cfg.PG.URL, sqlgorm.MaxPoolSize(cfg.PG.PoolMax))

	return client.DB, err
}

func provideRedisCache(cfg *config.Config) (*redis.Client, error) {
	c, err := redisCache.New(cfg.Redis.DSN)

	return c.Client, err
}

func provideLocalCache() (*bigcache.BigCache, error) {
	c, err := local.New()

	return c.Client, err
}

func provideUserRepo(sqlDatasource *user.DatasourceImpl, redisCacheDatasource *cache.RedisCacheDataSourceImpl, bigCacheDatasource *cache.BigCacheDataSourceImpl) *user.RepoImpl {
	return user.NewRepoImpl(sqlDatasource, redisCacheDatasource, bigCacheDatasource)
}

func provideRoleRepo(sqlDatasource *role.DatasourceImpl, redisCacheDatasource *cache.RedisCacheDataSourceImpl, bigCacheDatasource *cache.BigCacheDataSourceImpl) *role.RepoImpl {
	return role.NewRepoImpl(sqlDatasource, redisCacheDatasource, bigCacheDatasource)
}

func provideServices(user application_user.IService, role application_role.IService) v1.Services {
	return v1.Services{
		User: user,
		Role: role,
	}
}

func provideUserService(userRepo *user.RepoImpl, l pwlogger.Interface) application_user.IService {
	return application_user.NewServiceImpl(userRepo, l)
}

func provideRoleService(roleRepo *role.RepoImpl, l pwlogger.Interface) application_role.IService {
	return application_role.NewServiceImpl(roleRepo, l)
}

func provideHTTPServer(handler *gin.Engine, cfg *config.Config) *httpserver.Server {
	return httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}

var appSet = wire.NewSet(
	providePostgres,
	provideRedisCache,
	provideLocalCache,
	user.NewDatasourceImpl,
	role.NewDatasourceImpl,
	cache.NewRedisCacheDataSourceImpl,
	cache.NewBigCacheDataSourceImp,
	provideUserRepo,
	provideRoleRepo,
	provideUserService,
	provideRoleService,
	provideServices,
	v1.NewRouter,
	provideHTTPServer,
)

func NewHTTPServer(cfg *config.Config, l pwlogger.Interface) (*httpserver.Server, error) {
	wire.Build(appSet)

	return &httpserver.Server{}, nil
}
