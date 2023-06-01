package repo

import (
	"context"
	"fmt"

	entity "github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"gorm.io/gorm"
)

var _ datasource.DataSource[*entity.User] = (*CRUDDatasourceImpl[*entity.User])(nil)

// CRUDDatasourceImpl -.
type CRUDDatasourceImpl[T datasource.Entity] struct {
	DB *gorm.DB
}

// NewCRUDDatasourceImpl -.
func NewCRUDDatasourceImpl[T datasource.Entity](db *gorm.DB) *CRUDDatasourceImpl[T] {
	return &CRUDDatasourceImpl[T]{DB: db}
}

// GetByID -.
func (r *CRUDDatasourceImpl[T]) GetByID(ctx context.Context, model T) (T, error) {
	err := r.DB.First(model, model.GetID()).Error
	if err != nil {
		return nil, fmt.Errorf("CRUDDatasourceImpl - GetByID - r.DB.First: %w", err)
	}

	return model, nil
}

// Create -.
func (r *CRUDDatasourceImpl[T]) Create(ctx context.Context, model T) (T, error) {
	err := r.DB.Create(model).Error
	if err != nil {
		return nil, fmt.Errorf("CRUDDatasourceImpl - Create - r.DB.Create: %w", err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl[T]) Update(ctx context.Context, model T) (T, error) {
	err := r.DB.Save(model).Error
	if err != nil {
		return nil, fmt.Errorf("CRUDDatasourceImpl - Update - r.DB.Save: %w", err)
	}

	return model, nil
}

// Delete -.
func (r *CRUDDatasourceImpl[T]) Delete(ctx context.Context, model T) error {
	err := r.DB.Delete(model, model.GetID()).Error
	if err != nil {
		return fmt.Errorf("CRUDDatasourceImpl - Delete - r.DB.Delete: %w", err)
	}

	return nil
}
