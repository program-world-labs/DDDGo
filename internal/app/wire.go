//go:build wireinject
// +build wireinject

package app

import (
	"fmt"

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
	"github.com/program-world-labs/DDDGo/internal/infra/datasource/cache"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource/sql"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	"github.com/program-world-labs/DDDGo/internal/infra/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/role"
	"github.com/program-world-labs/DDDGo/internal/infra/user"
	"github.com/program-world-labs/DDDGo/pkg/cache/local"
	redisCache "github.com/program-world-labs/DDDGo/pkg/cache/redis"
	"github.com/program-world-labs/DDDGo/pkg/httpserver"
	"github.com/program-world-labs/DDDGo/pkg/pwsql"
)

func providePostgres(cfg *config.Config) (pwsql.ISQLGorm, error) {
	// postgres://user:password@localhost:5432/postgres
	port := fmt.Sprint(cfg.SQL.Port)
	dsn := cfg.SQL.Type + "://" + cfg.SQL.User + ":" + cfg.SQL.Password + "@" + cfg.SQL.Host + ":" + port + "/" + cfg.SQL.DB
	client, err := pwsql.New(dsn, pwsql.MaxPoolSize(cfg.SQL.PoolMax))
	client.GetDB().AutoMigrate(&dto.User{}, &dto.Role{})

	return client, err
}

func provideRedisCache(cfg *config.Config) (*redis.Client, error) {
	// redis://localhost:6379/0
	port := fmt.Sprint(cfg.Redis.Port)
	db := fmt.Sprint(cfg.Redis.DB)
	dsn := "redis://" + cfg.Redis.Host + ":" + port + "/" + db
	c, err := redisCache.New(dsn)

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

func provideTransactionRepo(datasource *sql.TransactionDataSourceImpl) *repository.TransactionRunRepoImpl {
	return repository.NewTransactionRunRepoImpl(datasource)
}

func provideUserRepo(sqlDatasource *sql.CRUDDatasourceImpl, bigCacheDatasource *cache.BigCacheDataSourceImpl, client *rockscache.Client) *user.RepoImpl {
	userCache := cache.NewRedisCacheDataSourceImpl(client, sqlDatasource)
	return user.NewRepoImpl(sqlDatasource, userCache, bigCacheDatasource)
}

func provideRoleRepo(sqlDatasource *sql.CRUDDatasourceImpl, bigCacheDatasource *cache.BigCacheDataSourceImpl, client *rockscache.Client) *role.RepoImpl {
	roleCache := cache.NewRedisCacheDataSourceImpl(client, sqlDatasource)
	return role.NewRepoImpl(sqlDatasource, roleCache, bigCacheDatasource)
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

func provideRoleService(roleRepo *role.RepoImpl, transactionRepo *repository.TransactionRunRepoImpl, l pwlogger.Interface) application_role.IService {
	return application_role.NewServiceImpl(roleRepo, transactionRepo, l)
}

func provideHTTPServer(handler *gin.Engine, cfg *config.Config) *httpserver.Server {
	return httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}

var appSet = wire.NewSet(
	providePostgres,
	provideRedisCache,
	provideLocalCache,
	provideRocksCache,
	sql.NewTransactionRunDataSourceImpl,
	sql.NewCRUDDatasourceImpl,
	cache.NewBigCacheDataSourceImp,
	provideTransactionRepo,
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
