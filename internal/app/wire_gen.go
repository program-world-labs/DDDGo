// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/allegro/bigcache/v3"
	"github.com/dtm-labs/rockscache"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/program-world-labs/DDDGo/config"
	"github.com/program-world-labs/DDDGo/internal/adapter/http/v1"
	message3 "github.com/program-world-labs/DDDGo/internal/adapter/message"
	"github.com/program-world-labs/DDDGo/internal/application"
	role2 "github.com/program-world-labs/DDDGo/internal/application/role"
	user2 "github.com/program-world-labs/DDDGo/internal/application/user"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource/cache"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource/sql"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	"github.com/program-world-labs/DDDGo/internal/infra/repository"
	"github.com/program-world-labs/DDDGo/internal/infra/role"
	"github.com/program-world-labs/DDDGo/internal/infra/user"
	"github.com/program-world-labs/DDDGo/pkg/cache/local"
	redis2 "github.com/program-world-labs/DDDGo/pkg/cache/redis"
	"github.com/program-world-labs/DDDGo/pkg/httpserver"
	message2 "github.com/program-world-labs/DDDGo/pkg/message"
	"github.com/program-world-labs/DDDGo/pkg/pwsql"
	"github.com/program-world-labs/pwlogger"
	"github.com/redis/go-redis/v9"
)

// Injectors from wire.go:

func NewHTTPServer(cfg *config.Config, l pwlogger.Interface) (*httpserver.Server, error) {
	isqlGorm, err := providePostgres(cfg)
	if err != nil {
		return nil, err
	}
	crudDatasourceImpl := sql.NewCRUDDatasourceImpl(isqlGorm)
	bigCache, err := provideLocalCache()
	if err != nil {
		return nil, err
	}
	bigCacheDataSourceImpl := cache.NewBigCacheDataSourceImp(bigCache)
	client, err := provideRedisCache(cfg)
	if err != nil {
		return nil, err
	}
	rockscacheClient := provideRocksCache(client)
	repoImpl := provideRoleRepo(crudDatasourceImpl, bigCacheDataSourceImpl, rockscacheClient)
	userRepoImpl := provideUserRepo(crudDatasourceImpl, bigCacheDataSourceImpl, rockscacheClient)
	transactionDataSourceImpl := sql.NewTransactionRunDataSourceImpl(isqlGorm)
	transactionRunRepoImpl := provideTransactionRepo(transactionDataSourceImpl)
	kafkaMessage, err := provideKafkaMessage(cfg)
	if err != nil {
		return nil, err
	}
	iService := provideUserService(repoImpl, userRepoImpl, transactionRunRepoImpl, kafkaMessage, l)
	roleIService := provideRoleService(repoImpl, userRepoImpl, transactionRunRepoImpl, kafkaMessage, l)
	services := provideServices(iService, roleIService)
	engine := v1.NewRouter(l, services, cfg)
	server := provideHTTPServer(engine, cfg)
	return server, nil
}

func NewMessageRouter(cfg *config.Config, l pwlogger.Interface) (*message.Router, error) {
	kafkaMessage, err := provideKafkaMessage(cfg)
	if err != nil {
		return nil, err
	}
	eventTypeMapper := provideEventTypeMapper()
	isqlGorm, err := providePostgres(cfg)
	if err != nil {
		return nil, err
	}
	crudDatasourceImpl := sql.NewCRUDDatasourceImpl(isqlGorm)
	bigCache, err := provideLocalCache()
	if err != nil {
		return nil, err
	}
	bigCacheDataSourceImpl := cache.NewBigCacheDataSourceImp(bigCache)
	client, err := provideRedisCache(cfg)
	if err != nil {
		return nil, err
	}
	rockscacheClient := provideRocksCache(client)
	repoImpl := provideRoleRepo(crudDatasourceImpl, bigCacheDataSourceImpl, rockscacheClient)
	userRepoImpl := provideUserRepo(crudDatasourceImpl, bigCacheDataSourceImpl, rockscacheClient)
	transactionDataSourceImpl := sql.NewTransactionRunDataSourceImpl(isqlGorm)
	transactionRunRepoImpl := provideTransactionRepo(transactionDataSourceImpl)
	iService := provideUserService(repoImpl, userRepoImpl, transactionRunRepoImpl, kafkaMessage, l)
	roleIService := provideRoleService(repoImpl, userRepoImpl, transactionRunRepoImpl, kafkaMessage, l)
	services := provideServices(iService, roleIService)
	router, err := provideMessageRouter(kafkaMessage, eventTypeMapper, services, l)
	if err != nil {
		return nil, err
	}
	return router, nil
}

// wire.go:

func providePostgres(cfg *config.Config) (pwsql.ISQLGorm, error) {

	port := fmt.Sprint(cfg.SQL.Port)
	dsn := cfg.SQL.Type + "://" + cfg.SQL.User + ":" + cfg.SQL.Password + "@" + cfg.SQL.Host + ":" + port + "/" + cfg.SQL.DB
	client, err := pwsql.New(dsn, pwsql.MaxPoolSize(cfg.SQL.PoolMax))
	client.GetDB().AutoMigrate(&dto.User{}, &dto.Role{})

	return client, err
}

func provideRedisCache(cfg *config.Config) (*redis.Client, error) {

	port := fmt.Sprint(cfg.Redis.Port)
	db := fmt.Sprint(cfg.Redis.DB)
	dsn := "redis://" + cfg.Redis.Host + ":" + port + "/" + db
	c, err := redis2.New(dsn)

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

func provideServices(user3 user2.IService, role3 role2.IService) application.Services {
	return application.Services{
		User: user3,
		Role: role3,
	}
}

func provideUserService(roleRepo *role.RepoImpl, userRepo *user.RepoImpl, transactionRepo *repository.TransactionRunRepoImpl, eventProducer *message2.KafkaMessage, l pwlogger.Interface) user2.IService {
	return user2.NewServiceImpl(roleRepo, userRepo, transactionRepo, eventProducer, l)
}

func provideRoleService(roleRepo *role.RepoImpl, userRepo *user.RepoImpl, transactionRepo *repository.TransactionRunRepoImpl, eventProducer *message2.KafkaMessage, l pwlogger.Interface) role2.IService {
	return role2.NewServiceImpl(roleRepo, userRepo, transactionRepo, eventProducer, l)
}

func provideHTTPServer(handler *gin.Engine, cfg *config.Config) *httpserver.Server {
	return httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}

func provideKafkaMessage(cfg *config.Config) (*message2.KafkaMessage, error) {
	return message2.NewKafkaMessage(cfg.Kafka.Brokers, cfg.Kafka.GroupID)
}

func provideMessageRouter(handler *message2.KafkaMessage, mapper *event.EventTypeMapper, s application.Services, l pwlogger.Interface) (*message.Router, error) {
	return message3.NewRouter(handler, mapper, s, l)
}

func provideEventTypeMapper() *event.EventTypeMapper {
	return event.NewEventTypeMapper()
}

var appSet = wire.NewSet(
	providePostgres,
	provideRedisCache,
	provideLocalCache,
	provideRocksCache, sql.NewTransactionRunDataSourceImpl, sql.NewCRUDDatasourceImpl, cache.NewBigCacheDataSourceImp, provideTransactionRepo,
	provideUserRepo,
	provideRoleRepo,
	provideKafkaMessage,
	provideMessageRouter,
	provideEventTypeMapper,
	provideUserService,
	provideRoleService,
	provideServices, v1.NewRouter, provideHTTPServer,
)
