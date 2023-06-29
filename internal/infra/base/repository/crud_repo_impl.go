package repository

import (
	"context"
	"fmt"

	"github.com/jinzhu/copier"
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
	info, err := Transform(e, r.DTOEntity)
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - GetByID - r.DTOEntity.Transform: %w", err)
	}

	// 先從Local Cache取得資料
	data, err := r.Cache.Get(ctx, info)
	if err == nil {
		return data, nil
	}

	// 若Local Cache沒有資料，則從Redis取得資料
	data, err = r.Redis.Get(ctx, info)
	if err == nil {
		return data, nil
	}

	// 若Redis沒有資料，則從DB取得資料
	data, err = r.DB.GetByID(ctx, info)
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - GetByID - r.DB.GetByID: %w", err)
	}

	return data, nil
}

// Create -.
func (r *CRUDImpl) Create(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	err := copier.Copy(r.DTOEntity, e.Self())
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - Create - r.DTOEntity.Transform: %w", err)
	}

	_, err = r.DB.Create(ctx, r.DTOEntity)
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - Create - r.DB.Create: %w", err)
	}

	return e, nil
}

// Update -.
func (r *CRUDImpl) Update(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	info, err := Transform(e, r.DTOEntity)
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - Update - r.DTOEntity.Transform: %w", err)
	}

	_, err = r.DB.Update(ctx, info)
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - Update - r.DB.Update: %w", err)
	}

	return e, nil
}

// Delete -.
func (r *CRUDImpl) Delete(ctx context.Context, e domain.IEntity) error {
	info, err := Transform(e, r.DTOEntity)
	if err != nil {
		return fmt.Errorf("CRUDImpl - Delete - r.DTOEntity.Transform: %w", err)
	}

	err = r.DB.Delete(ctx, info)
	if err != nil {
		return fmt.Errorf("CRUDImpl - Delete - r.DB.Delete: %w", err)
	}

	return nil
}
