package sql

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/infra/base/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
	"github.com/program-world-labs/DDDGo/pkg/pwsql"
)

var _ datasource.IDataSource = (*CRUDDatasourceImpl)(nil)

// CRUDDatasourceImpl -.
type CRUDDatasourceImpl struct {
	DB *gorm.DB
}

// NewCRUDDatasourceImpl -.
func NewCRUDDatasourceImpl(db pwsql.ISQLGorm) *CRUDDatasourceImpl {
	return &CRUDDatasourceImpl{DB: db.GetDB()}
}

// GetByID -.
func (r *CRUDDatasourceImpl) GetByID(_ context.Context, model entity.IEntity) (entity.IEntity, error) {
	err := r.DB.First(model, model.GetID()).Error
	if err != nil {
		return nil, fmt.Errorf("CRUDDatasourceImpl - GetByID - r.DB.First: %w", err)
	}

	return model, nil
}

// Create -.
func (r *CRUDDatasourceImpl) Create(_ context.Context, model entity.IEntity) (entity.IEntity, error) {
	err := r.DB.Create(model).Error
	if err != nil {
		return nil, fmt.Errorf("CRUDDatasourceImpl - Create - r.DB.Create: %w", err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl) Update(_ context.Context, model entity.IEntity) (entity.IEntity, error) {
	err := r.DB.Save(model).Error
	if err != nil {
		return nil, fmt.Errorf("CRUDDatasourceImpl - Update - r.DB.Save: %w", err)
	}

	return model, nil
}

// UpdateWithFields -.
func (r *CRUDDatasourceImpl) UpdateWithFields(_ context.Context, model entity.IEntity, fields []string) (entity.IEntity, error) {
	err := r.DB.Model(model).Select(fields).Updates(model).Error
	if err != nil {
		return nil, fmt.Errorf("CRUDDatasourceImpl - UpdateWithFields - r.DB.Model: %w", err)
	}

	return model, nil
}

// Delete -.
func (r *CRUDDatasourceImpl) Delete(_ context.Context, model entity.IEntity) error {
	err := r.DB.Delete(model, model.GetID()).Error
	if err != nil {
		return fmt.Errorf("CRUDDatasourceImpl - Delete - r.DB.Delete: %w", err)
	}

	return nil
}
