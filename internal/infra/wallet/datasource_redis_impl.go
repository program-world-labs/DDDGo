package wallet

import (
	"github.com/dtm-labs/rockscache"

	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource/cache"
)

var _ datasource.ICacheDataSource = (*CacheDatasourceImpl)(nil)

// DatasourceImpl -.
type CacheDatasourceImpl struct {
	cache.RedisCacheDataSourceImpl
}

// NewDatasourceImpl -.
func NewCacheDatasourceImpl(c *rockscache.Client, sqlDatasource datasource.IDataSource) *CacheDatasourceImpl {
	return &CacheDatasourceImpl{RedisCacheDataSourceImpl: *cache.NewRedisCacheDataSourceImpl(c, sqlDatasource)}
}
