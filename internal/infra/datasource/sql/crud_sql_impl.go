package sql

import (
	"context"

	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
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
func (r *CRUDDatasourceImpl) GetByID(_ context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	err := r.DB.First(model, model.GetID()).Error
	if err != nil {
		return nil, NewGetError(err)
	}

	return model, nil
}

// Create -.
func (r *CRUDDatasourceImpl) Create(_ context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	err := r.DB.Create(model).Error
	if err != nil {
		return nil, NewCreateError(err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl) Update(_ context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	err := r.DB.Save(model).Error
	if err != nil {
		return nil, NewUpdateError(err)
	}

	return model, nil
}

// UpdateWithFields -.
func (r *CRUDDatasourceImpl) UpdateWithFields(_ context.Context, model dto.IRepoEntity, fields []string) (dto.IRepoEntity, error) {
	err := r.DB.Model(model).Select(fields).Updates(model).Error
	if err != nil {
		return nil, NewUpdateWithFieldsError(err)
	}

	return model, nil
}

// Delete -.
func (r *CRUDDatasourceImpl) Delete(_ context.Context, model dto.IRepoEntity) error {
	err := r.DB.Delete(model, model.GetID()).Error
	if err != nil {
		return NewDeleteError(err)
	}

	return nil
}

func (r *CRUDDatasourceImpl) GetAll(ctx context.Context, e dto.IRepoEntity, sq *domain.SearchQuery) ([]dto.IRepoEntity, error) {
	var data []dto.IRepoEntity
	err := r.DB.Where(sq.GetWhere(), sq.GetArgs()...).Order(sq.GetOrder()).Find(&data).Error

	if err != nil {
		return nil, NewGetAllError(err)
	}

	return data, nil
}

// Create -.
func (r *CRUDDatasourceImpl) CreateTx(_ context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return nil, NewCastError(ErrCastToEntityFailed)
	}

	err := t.Create(model).Error
	if err != nil {
		return nil, NewCreateError(err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl) UpdateTx(_ context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return nil, NewCastError(ErrCastToEntityFailed)
	}

	err := t.Save(model).Error
	if err != nil {
		return nil, NewUpdateError(err)
	}

	return model, nil
}

// UpdateWithFields -.
func (r *CRUDDatasourceImpl) UpdateWithFieldsTx(_ context.Context, model dto.IRepoEntity, fields []string, tx domain.ITransactionEvent) error {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return NewCastError(ErrCastToEntityFailed)
	}

	err := t.Model(model).Select(fields).Updates(model).Error
	if err != nil {
		return NewUpdateWithFieldsError(err)
	}

	return nil
}

// Delete -.
func (r *CRUDDatasourceImpl) DeleteTx(_ context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) error {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return NewCastError(ErrCastToEntityFailed)
	}

	err := t.Delete(model, model.GetID()).Error
	if err != nil {
		return NewDeleteError(err)
	}

	return nil
}
