package repository

import (
	"context"
	"fmt"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
)

// CRUDImpl -.
type CRUDImpl struct {
	DB    datasource.IDataSource
	Redis datasource.ICacheDataSource
	Cache datasource.ICacheDataSource
}

// NewCRUDImpl -.
func NewCRUDImpl(db datasource.IDataSource, redis datasource.ICacheDataSource, cache datasource.ICacheDataSource) *CRUDImpl {
	return &CRUDImpl{DB: db, Redis: redis, Cache: cache}
}

// GetByID -.
func (r *CRUDImpl) GetByID(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	// 先從Local Cache取得資料
	data, err := r.Cache.Get(ctx, e)
	if err == nil {
		return data, nil
	}

	// 若Local Cache沒有資料，則從Redis取得資料
	data, err = r.Redis.Get(ctx, e)
	if err == nil {
		return data, nil
	}

	// 若Redis沒有資料，則從DB取得資料
	data, err = r.DB.GetByID(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - GetByID - r.DB.GetByID: %w", err)
	}

	return data, nil
}

// Create -.
func (r *CRUDImpl) Create(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	_, err := r.DB.Create(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - Create - r.DB.Create: %w", err)
	}

	return e, nil
}

// Update -.
func (r *CRUDImpl) Update(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	_, err := r.DB.Update(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - Update - r.DB.Update: %w", err)
	}

	return e, nil
}

// Delete -.
func (r *CRUDImpl) Delete(ctx context.Context, e domain.IEntity) error {
	err := r.DB.Delete(ctx, e)
	if err != nil {
		return fmt.Errorf("CRUDImpl - Delete - r.DB.Delete: %w", err)
	}

	return nil
}
