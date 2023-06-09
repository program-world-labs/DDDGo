package sql

import (
	"context"

	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	"github.com/program-world-labs/DDDGo/pkg/pwsql"
)

var _ datasource.IRelationDataSource = (*CRUDDatasourceImpl)(nil)

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
	// 加入預載入
	if len(model.GetPreloads()) > 0 {
		for _, preload := range model.GetPreloads() {
			r.DB = r.DB.Preload(preload)
		}
	}

	err := r.DB.Table(model.TableName()).First(&model, "id = ?", model.GetID()).Error

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLGet, err)
	}

	return model, nil
}

// Create -.
func (r *CRUDDatasourceImpl) Create(_ context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	err := r.DB.Create(model).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLCreate, err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl) Update(_ context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	err := r.DB.Save(model).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	return model, nil
}

// UpdateWithFields -.
func (r *CRUDDatasourceImpl) UpdateWithFields(_ context.Context, model dto.IRepoEntity, fields []string) error {
	err := r.DB.Model(model).Select(fields).Updates(model).Error
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLUpdateWithFields, err)
	}

	return nil
}

// Delete -.
func (r *CRUDDatasourceImpl) Delete(_ context.Context, model dto.IRepoEntity) error {
	err := r.DB.Delete(model, model.GetID()).Error
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLDelete, err)
	}

	return nil
}

func (r *CRUDDatasourceImpl) GetAll(_ context.Context, sq *domain.SearchQuery, model dto.IRepoEntity) (*dto.List, error) {
	// 加入預載入
	if len(model.GetPreloads()) > 0 {
		for _, preload := range model.GetPreloads() {
			r.DB = r.DB.Preload(preload)
		}
	}

	var data = model.GetListType()

	err := r.DB.Table(model.TableName()).Limit(sq.Page.Limit).Offset(sq.Page.Offset).Where(sq.GetWhere(), sq.GetArgs()...).Order(sq.GetOrder()).Find(&data).Error

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLGetAll, err)
	}

	var count int64
	err = r.DB.Table(model.TableName()).Where(sq.GetWhere(), sq.GetArgs()...).Count(&count).Error

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLGetAll, err)
	}

	// var list []dto.IRepoEntity

	// for _, item := range data {
	// 	list = append(list, item.(dto.IRepoEntity))
	// }

	var result = &dto.List{
		Limit:  sq.Page.Limit,
		Offset: sq.Page.Offset,
		Total:  int(count),
		Data:   data,
	}

	return result, nil
}

// Create -.
func (r *CRUDDatasourceImpl) CreateTx(_ context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.Create(model).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLCreate, err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl) UpdateTx(_ context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.Save(model).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	return model, nil
}

// UpdateWithFields -.
func (r *CRUDDatasourceImpl) UpdateWithFieldsTx(_ context.Context, model dto.IRepoEntity, fields []string, tx domain.ITransactionEvent) error {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.Model(model).Select(fields).Updates(model).Error
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLUpdateWithFields, err)
	}

	return nil
}

// Delete -.
func (r *CRUDDatasourceImpl) DeleteTx(_ context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) error {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.Delete(model, model.GetID()).Error
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLDelete, err)
	}

	return nil
}

// AppendAssociation -.
func (r *CRUDDatasourceImpl) AppendAssociation(_ context.Context, key string, model dto.IRepoEntity, appendModel []dto.IRepoEntity) error {
	err := r.DB.Model(model).Association(key).Append(appendModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLAppendAssociation, err)
	}

	return nil
}

// ReplaceAssociation -.
func (r *CRUDDatasourceImpl) ReplaceAssociation(_ context.Context, key string, model dto.IRepoEntity, replaceModel []dto.IRepoEntity) error {
	err := r.DB.Model(model).Association(key).Replace(replaceModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLReplaceAssociation, err)
	}

	return nil
}

// RemoveAssociation -.
func (r *CRUDDatasourceImpl) RemoveAssociation(_ context.Context, key string, model dto.IRepoEntity, removeModel []dto.IRepoEntity) error {
	err := r.DB.Model(model).Association(key).Delete(removeModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLRemoveAssociation, err)
	}

	return nil
}

// GetAssociationCount -.
func (r *CRUDDatasourceImpl) GetAssociationCount(_ context.Context, key string, model dto.IRepoEntity) (int64, error) {
	count := r.DB.Model(model).Association(key).Count()

	return count, nil
}

// AppendAssociationTx -.
func (r *CRUDDatasourceImpl) AppendAssociationTx(_ context.Context, key string, model dto.IRepoEntity, appendModel []dto.IRepoEntity, tx domain.ITransactionEvent) error {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.Model(model).Association(key).Append(appendModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLAppendAssociation, err)
	}

	return nil
}

// ReplaceAssociationTx -.
func (r *CRUDDatasourceImpl) ReplaceAssociationTx(_ context.Context, key string, model dto.IRepoEntity, replaceModel []dto.IRepoEntity, tx domain.ITransactionEvent) error {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.Model(model).Association(key).Replace(replaceModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLReplaceAssociation, err)
	}

	return nil
}

// RemoveAssociationTx -.
func (r *CRUDDatasourceImpl) RemoveAssociationTx(_ context.Context, key string, model dto.IRepoEntity, removeModel []dto.IRepoEntity, tx domain.ITransactionEvent) error {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.Model(model).Association(key).Delete(removeModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLRemoveAssociation, err)
	}

	return nil
}
