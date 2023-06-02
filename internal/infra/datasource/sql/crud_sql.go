package sql

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	entity "github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
)

var _ datasource.IDataSource[*entity.User] = (*CRUDDatasourceImpl[*entity.User])(nil)

// CRUDDatasourceImpl -.
type CRUDDatasourceImpl[T datasource.IEntity] struct {
	DB *gorm.DB
}

// NewCRUDDatasourceImpl -.
func NewCRUDDatasourceImpl[T datasource.IEntity](db *gorm.DB) *CRUDDatasourceImpl[T] {
	return &CRUDDatasourceImpl[T]{DB: db}
}

// GetByID -.
func (r *CRUDDatasourceImpl[T]) GetByID(_ context.Context, model T) (T, error) {
	err := r.DB.First(model, model.GetID()).Error
	if err != nil {
		return nil, fmt.Errorf("CRUDDatasourceImpl - GetByID - r.DB.First: %w", err)
	}

	return model, nil
}

// Create -.
func (r *CRUDDatasourceImpl[T]) Create(_ context.Context, model T) (T, error) {
	err := r.DB.Create(model).Error
	if err != nil {
		return nil, fmt.Errorf("CRUDDatasourceImpl - Create - r.DB.Create: %w", err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl[T]) Update(_ context.Context, model T) (T, error) {
	err := r.DB.Save(model).Error
	if err != nil {
		return nil, fmt.Errorf("CRUDDatasourceImpl - Update - r.DB.Save: %w", err)
	}

	return model, nil
}

// Delete -.
func (r *CRUDDatasourceImpl[T]) Delete(_ context.Context, model T) error {
	err := r.DB.Delete(model, model.GetID()).Error
	if err != nil {
		return fmt.Errorf("CRUDDatasourceImpl - Delete - r.DB.Delete: %w", err)
	}

	return nil
}
