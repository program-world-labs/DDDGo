//go:build wireinject
// +build wireinject

package app

import (
	"github.com/allegro/bigcache/v3"
	"github.com/dtm-labs/rockscache"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/program-world-labs/pwlogger"
	"github.com/redis/go-redis/v9"

	"github.com/program-world-labs/DDDGo/config"
	v1 "github.com/program-world-labs/DDDGo/internal/adapter/http/v1"
	application_role "github.com/program-world-labs/DDDGo/internal/application/role"
	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource/cache"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	"github.com/program-world-labs/DDDGo/internal/infra/role"
	"github.com/program-world-labs/DDDGo/internal/infra/user"
	"github.com/program-world-labs/DDDGo/pkg/cache/local"
	redisCache "github.com/program-world-labs/DDDGo/pkg/cache/redis"
	"github.com/program-world-labs/DDDGo/pkg/httpserver"
	"github.com/program-world-labs/DDDGo/pkg/pwsql"
)

func providePostgres(cfg *config.Config) (pwsql.ISQLGorm, error) {
	client, err := pwsql.New(cfg.PG.URL, pwsql.MaxPoolSize(cfg.PG.PoolMax))
	client.GetDB().AutoMigrate(&dto.User{}, &dto.Role{})

	return client, err
}

func provideRedisCache(cfg *config.Config) (*redis.Client, error) {
	c, err := redisCache.New(cfg.Redis.DSN)

	return c.Client, err
}

func provideRocksCache(r *redis.Client) *rockscache.Client {
	rc := rockscache.NewClient(r, rockscache.NewDefaultOptions())

	return rc
}

func provideLocalCache() (*bigcache.BigCache, error) {
	c, err := local.New()

	return c.Client, err
}

func provideUserCacheDatasource(client *rockscache.Client, sqlDatasource *user.DatasourceImpl) *user.CacheDatasourceImpl {
	return user.NewCacheDatasourceImpl(client, sqlDatasource)
}

func provideRoleCacheDatasource(client *rockscache.Client, sqlDatasource *role.DatasourceImpl) *role.CacheDatasourceImpl {
	return role.NewCacheDatasourceImpl(client, sqlDatasource)
}

func provideUserRepo(sqlDatasource *user.DatasourceImpl, redisCacheDatasource *user.CacheDatasourceImpl, bigCacheDatasource *cache.BigCacheDataSourceImpl) *user.RepoImpl {
	return user.NewRepoImpl(sqlDatasource, redisCacheDatasource, bigCacheDatasource)
}

func provideRoleRepo(sqlDatasource *role.DatasourceImpl, redisCacheDatasource *role.CacheDatasourceImpl, bigCacheDatasource *cache.BigCacheDataSourceImpl) *role.RepoImpl {
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
	provideRocksCache,
	user.NewDatasourceImpl,
	role.NewDatasourceImpl,
	provideUserCacheDatasource,
	provideRoleCacheDatasource,
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
