package repository

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

var _ domain.ICRUDRepository = (*CRUDImpl)(nil)

// CRUDImpl -.
type CRUDImpl struct {
	DB        datasource.IDataSource
	Redis     datasource.ICacheDataSource
	Cache     datasource.ICacheDataSource
	DTOEntity entity.IEntity
}

// NewCRUDImpl -.
func NewCRUDImpl(db datasource.IDataSource, redis datasource.ICacheDataSource, cache datasource.ICacheDataSource, data entity.IEntity) *CRUDImpl {
	return &CRUDImpl{DB: db, Redis: redis, Cache: cache, DTOEntity: data}
}

// GetByID -.
func (r *CRUDImpl) GetByID(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	// 先從Local Cache取得資料
	data, err := r.Cache.Get(ctx, info)
	if err == nil {
		return data, nil
	}

	// 從Redis取得資料, 取不到資料會自動從db取得資料
	data, err = r.Redis.Get(ctx, info)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	// 將資料存入Local Cache
	_, err = r.Cache.Set(ctx, data)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	return data, nil
}

// Create -.
func (r *CRUDImpl) Create(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	_, err = r.DB.Create(ctx, info)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	return e, nil
}

// Update -.
func (r *CRUDImpl) Update(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	_, err = r.DB.Update(ctx, info)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	return e, nil
}

// UpdateWithFields -.
func (r *CRUDImpl) UpdateWithFields(ctx context.Context, e domain.IEntity, keys []string) (domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	_, err = r.DB.UpdateWithFields(ctx, info, keys)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	return e, nil
}

// Delete -.
func (r *CRUDImpl) Delete(ctx context.Context, e domain.IEntity) error {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return NewDatasourceError(err)
	}

	err = r.DB.Delete(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}

	return nil
}
