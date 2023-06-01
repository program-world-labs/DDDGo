package repository

import (
	"context"
	"fmt"

	"gitlab.com/demojira/template.git/internal/infra/datasource"
)

// CRUDImpl -.
type CRUDImpl[T datasource.Entity] struct {
	DB    datasource.DataSource[T]
	Redis datasource.CacheDataSource[T]
	Cache datasource.CacheDataSource[T]
}

// NewCRUDImpl -.
func NewCRUDImpl[T datasource.Entity](db datasource.DataSource[T], redis datasource.CacheDataSource[T], cache datasource.CacheDataSource[T]) *CRUDImpl[T] {
	return &CRUDImpl[T]{DB: db, Redis: redis, Cache: cache}
}

// GetByID -.
func (r *CRUDImpl[T]) GetByID(ctx context.Context, e T) (T, error) {
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
func (r *CRUDImpl[T]) Create(ctx context.Context, e T) (T, error) {
	_, err := r.DB.Create(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - Create - r.DB.Create: %w", err)
	}
	return e, nil
}

// Update -.
func (r *CRUDImpl[T]) Update(ctx context.Context, e T) (T, error) {
	_, err := r.DB.Update(ctx, e)
	if err != nil {
		return nil, fmt.Errorf("CRUDImpl - Update - r.DB.Update: %w", err)
	}
	return e, nil
}

// Delete -.
func (r *CRUDImpl[T]) Delete(ctx context.Context, e T) error {
	err := r.DB.Delete(ctx, e)
	if err != nil {
		return fmt.Errorf("CRUDImpl - Delete - r.DB.Delete: %w", err)
	}

	return nil
}
