//go:build wireinject
// +build wireinject

package app

import (
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/allegro/bigcache/v3"
	"github.com/dtm-labs/rockscache"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/program-world-labs/pwlogger"
	"github.com/redis/go-redis/v9"

	"github.com/program-world-labs/DDDGo/config"
	v1 "github.com/program-world-labs/DDDGo/internal/adapter/http/v1"
	adapter_message "github.com/program-world-labs/DDDGo/internal/adapter/message"
	"github.com/program-world-labs/DDDGo/internal/application"
	application_currency "github.com/program-world-labs/DDDGo/internal/application/currency"
	application_group "github.com/program-world-labs/DDDGo/internal/application/group"
	application_role "github.com/program-world-labs/DDDGo/internal/application/role"
	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
	application_wallet "github.com/program-world-labs/DDDGo/internal/application/wallet"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
	"github.com/program-world-labs/DDDGo/internal/infra/currency"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource/cache"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource/sql"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	"github.com/program-world-labs/DDDGo/internal/infra/group"
	"github.com/program-world-labs/DDDGo/internal/infra/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/role"
	"github.com/program-world-labs/DDDGo/internal/infra/user"
	"github.com/program-world-labs/DDDGo/internal/infra/wallet"
	"github.com/program-world-labs/DDDGo/pkg/cache/local"
	redisCache "github.com/program-world-labs/DDDGo/pkg/cache/redis"
	"github.com/program-world-labs/DDDGo/pkg/httpserver"
	pkg_message "github.com/program-world-labs/DDDGo/pkg/message"
	"github.com/program-world-labs/DDDGo/pkg/pwsql"
)

func provideSQL(cfg *config.Config) (pwsql.ISQLGorm, error) {
	// postgres://user:password@localhost:5432/postgres
	port := fmt.Sprint(cfg.SQL.Port)
	var dsn string
	switch cfg.SQL.Type {
	case "mysql":
		dsn = cfg.SQL.User + ":" + cfg.SQL.Password + "@tcp(" + cfg.SQL.Host + ":" + port + ")/" + cfg.SQL.DB + "?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true"
	case "postgresql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.SQL.Host, port, cfg.SQL.User, cfg.SQL.Password, cfg.SQL.DB)
	default:
	}
	client, err := pwsql.InitSQL(cfg.SQL.Type, dsn, cfg.SQL.PoolMax, cfg.SQL.ConnAttempts, cfg.SQL.ConnTimeout*time.Second)
	client.GetDB().AutoMigrate(&dto.User{}, &dto.Role{}, &dto.Group{}, &dto.Wallet{}, &dto.Currency{})

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
	rc.Options.StrongConsistency = true
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

func provideGroupRepo(sqlDatasource *sql.CRUDDatasourceImpl, bigCacheDatasource *cache.BigCacheDataSourceImpl, client *rockscache.Client) *group.RepoImpl {
	groupCache := cache.NewRedisCacheDataSourceImpl(client, sqlDatasource)
	return group.NewRepoImpl(sqlDatasource, groupCache, bigCacheDatasource)
}

func provideWalletRepo(sqlDatasource *sql.CRUDDatasourceImpl, bigCacheDatasource *cache.BigCacheDataSourceImpl, client *rockscache.Client) *wallet.RepoImpl {
	walletCache := cache.NewRedisCacheDataSourceImpl(client, sqlDatasource)
	return wallet.NewRepoImpl(sqlDatasource, walletCache, bigCacheDatasource)
}

func provideCurrencyRepo(sqlDatasource *sql.CRUDDatasourceImpl, bigCacheDatasource *cache.BigCacheDataSourceImpl, client *rockscache.Client) *currency.RepoImpl {
	currencyCache := cache.NewRedisCacheDataSourceImpl(client, sqlDatasource)
	return currency.NewRepoImpl(sqlDatasource, currencyCache, bigCacheDatasource)
}

func provideServices(user application_user.IService, role application_role.IService, group application_group.IService, wallet application_wallet.IService, currency application_currency.IService) application.Services {
	return application.Services{
		User:     user,
		Role:     role,
		Group:    group,
		Wallet:   wallet,
		Currency: currency,
	}
}

func provideUserService(roleRepo *role.RepoImpl, userRepo *user.RepoImpl, transactionRepo *repository.TransactionRunRepoImpl, eventProducer *pkg_message.KafkaMessage, l pwlogger.Interface) application_user.IService {
	return application_user.NewServiceImpl(roleRepo, userRepo, transactionRepo, eventProducer, l)
}

func provideRoleService(roleRepo *role.RepoImpl, userRepo *user.RepoImpl, transactionRepo *repository.TransactionRunRepoImpl, eventProducer *pkg_message.KafkaMessage, l pwlogger.Interface) application_role.IService {
	return application_role.NewServiceImpl(roleRepo, userRepo, transactionRepo, eventProducer, l)
}

func provideGroupService(groupRepo *group.RepoImpl, userRepo *user.RepoImpl, transactionRepo *repository.TransactionRunRepoImpl, eventProducer *pkg_message.KafkaMessage, l pwlogger.Interface) application_group.IService {
	return application_group.NewServiceImpl(groupRepo, userRepo, transactionRepo, eventProducer, l)
}

func provideWalletService(walletRepo *wallet.RepoImpl, userRepo *user.RepoImpl, transactionRepo *repository.TransactionRunRepoImpl, eventProducer *pkg_message.KafkaMessage, l pwlogger.Interface) application_wallet.IService {
	return application_wallet.NewServiceImpl(walletRepo, userRepo, transactionRepo, eventProducer, l)
}

func provideCurrencyService(currencyRepo *currency.RepoImpl, userRepo *user.RepoImpl, transactionRepo *repository.TransactionRunRepoImpl, eventProducer *pkg_message.KafkaMessage, l pwlogger.Interface) application_currency.IService {
	return application_currency.NewServiceImpl(currencyRepo, userRepo, transactionRepo, eventProducer, l)
}

func provideHTTPServer(handler *gin.Engine, cfg *config.Config) *httpserver.Server {
	return httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}

func provideKafkaMessage(cfg *config.Config) (*pkg_message.KafkaMessage, error) {
	return pkg_message.NewKafkaMessage(cfg.Kafka.Brokers, cfg.Kafka.GroupID)
}

func provideMessageRouter(handler *pkg_message.KafkaMessage, mapper *event.TypeMapper, s application.Services, l pwlogger.Interface) (*message.Router, error) {
	return adapter_message.NewRouter(handler, mapper, s, l)
}

func provideEventTypeMapper() *event.TypeMapper {
	return event.NewEventTypeMapper()
}

var appSet = wire.NewSet(
	provideSQL,
	provideRedisCache,
	provideLocalCache,
	provideRocksCache,
	sql.NewTransactionRunDataSourceImpl,
	sql.NewCRUDDatasourceImpl,
	cache.NewBigCacheDataSourceImp,
	provideTransactionRepo,
	provideUserRepo,
	provideRoleRepo,
	provideGroupRepo,
	provideWalletRepo,
	provideCurrencyRepo,
	provideKafkaMessage,
	provideMessageRouter,
	provideEventTypeMapper,
	provideUserService,
	provideRoleService,
	provideGroupService,
	provideWalletService,
	provideCurrencyService,
	provideServices,
	v1.NewRouter,
	provideHTTPServer,
)

func NewHTTPServer(cfg *config.Config, l pwlogger.Interface) (*httpserver.Server, error) {
	wire.Build(appSet)

	return &httpserver.Server{}, nil
}

// func NewMessageServer(cfg *config.Config, l pwlogger.Interface) (*pkg_message.KafkaMessage, error) {
// 	wire.Build(appSet)

// 	return &pkg_message.KafkaMessage{}, nil
// }

func NewMessageRouter(cfg *config.Config, l pwlogger.Interface) (*message.Router, error) {
	wire.Build(appSet)

	return &message.Router{}, nil
}
