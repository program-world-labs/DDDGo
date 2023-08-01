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
func (r *CRUDDatasourceImpl) GetByID(ctx context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	db := r.DB.WithContext(ctx).Table(model.TableName())

	//加入預載入
	if len(model.GetPreloads()) > 0 {
		for _, preload := range model.GetPreloads() {
			db = db.Preload(preload)
		}
	}

	err := db.First(&model, "id = ?", model.GetID()).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLGet, err)
	}

	return model, nil
}

// Create -.
func (r *CRUDDatasourceImpl) Create(ctx context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	// exec hook BeforeCreate
	if err := model.BeforeCreate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLCreate, err)
	}

	err := r.DB.WithContext(ctx).Create(model).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLCreate, err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl) Update(ctx context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	// exec hook BeforeUpdate
	if err := model.BeforeUpdate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	err := r.DB.WithContext(ctx).Omit("created_at").Omit("deleted_at").Updates(model).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	return model, nil
}

// UpdateWithFields -.
func (r *CRUDDatasourceImpl) UpdateWithFields(ctx context.Context, model dto.IRepoEntity, fields []string) (dto.IRepoEntity, error) {
	// exec hook BeforeUpdate
	if err := model.BeforeUpdate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	err := r.DB.WithContext(ctx).Model(model).Select(fields).Updates(model).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdateWithFields, err)
	}

	err = r.DB.WithContext(ctx).First(model, "id = ?", model.GetID()).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdateWithFields, err)
	}

	return model, nil
}

// Delete -.
func (r *CRUDDatasourceImpl) Delete(ctx context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	err := r.DB.WithContext(ctx).Delete(model, "id = ?", model.GetID()).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLDelete, err)
	}

	err = r.DB.WithContext(ctx).Unscoped().First(model, "id = ?", model.GetID()).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLDelete, err)
	}

	return model, nil
}

func (r *CRUDDatasourceImpl) GetAll(ctx context.Context, sq *domain.SearchQuery, model dto.IRepoEntity) (*dto.List, error) {
	db := r.DB.WithContext(ctx).Table(model.TableName()).Limit(sq.Page.Limit).Offset(sq.Page.Offset).Where(sq.GetWhere(), sq.GetArgs()...).Order(sq.GetOrder())

	if len(model.GetPreloads()) > 0 {
		for _, preload := range model.GetPreloads() {
			db = db.Preload(preload)
		}
	}

	var count int64

	dbCount := r.DB.WithContext(ctx).Table(model.TableName()).Where(sq.GetWhere(), sq.GetArgs()...).Where("deleted_at IS NULL")
	err := dbCount.Count(&count).Error

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLGetAll, err)
	}

	var data = model.GetListType()
	err = db.Find(&data).Error

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
func (r *CRUDDatasourceImpl) CreateTx(ctx context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	// exec hook BeforeCreate
	if err := model.BeforeCreate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLCreate, err)
	}

	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.WithContext(ctx).Create(model).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLCreate, err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl) UpdateTx(ctx context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	// exec hook BeforeUpdate
	if err := model.BeforeUpdate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.WithContext(ctx).Save(model).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	return model, nil
}

// UpdateWithFields -.
func (r *CRUDDatasourceImpl) UpdateWithFieldsTx(ctx context.Context, model dto.IRepoEntity, fields []string, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	// exec hook BeforeUpdate
	if err := model.BeforeUpdate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.WithContext(ctx).Model(model).Select(fields).Updates(model).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdateWithFields, err)
	}

	err = r.DB.WithContext(ctx).First(model, "id = ?", model.GetID()).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdateWithFields, err)
	}

	return model, nil
}

// Delete -.
func (r *CRUDDatasourceImpl) DeleteTx(ctx context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.WithContext(ctx).Delete(model, model.GetID()).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLDelete, err)
	}

	err = r.DB.WithContext(ctx).First(model, "id = ?", model.GetID()).Error
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLDelete, err)
	}

	return model, nil
}

// AppendAssociation -.
func (r *CRUDDatasourceImpl) AppendAssociation(ctx context.Context, key string, model dto.IRepoEntity, appendModel []dto.IRepoEntity) error {
	err := r.DB.WithContext(ctx).Model(model).Association(key).Append(appendModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLAppendAssociation, err)
	}

	return nil
}

// ReplaceAssociation -.
func (r *CRUDDatasourceImpl) ReplaceAssociation(ctx context.Context, key string, model dto.IRepoEntity, replaceModel []dto.IRepoEntity) error {
	err := r.DB.WithContext(ctx).Model(model).Association(key).Replace(replaceModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLReplaceAssociation, err)
	}

	return nil
}

// RemoveAssociation -.
func (r *CRUDDatasourceImpl) RemoveAssociation(ctx context.Context, key string, model dto.IRepoEntity, removeModel []dto.IRepoEntity) error {
	err := r.DB.WithContext(ctx).Model(model).Association(key).Delete(removeModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLRemoveAssociation, err)
	}

	return nil
}

// GetAssociationCount -.
func (r *CRUDDatasourceImpl) GetAssociationCount(ctx context.Context, key string, model dto.IRepoEntity) (int64, error) {
	count := r.DB.WithContext(ctx).Model(model).Association(key).Count()

	return count, nil
}

// AppendAssociationTx -.
func (r *CRUDDatasourceImpl) AppendAssociationTx(ctx context.Context, key string, model dto.IRepoEntity, appendModel []dto.IRepoEntity, tx domain.ITransactionEvent) error {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.WithContext(ctx).Model(model).Association(key).Append(appendModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLAppendAssociation, err)
	}

	return nil
}

// ReplaceAssociationTx -.
func (r *CRUDDatasourceImpl) ReplaceAssociationTx(ctx context.Context, key string, model dto.IRepoEntity, replaceModel []dto.IRepoEntity, tx domain.ITransactionEvent) error {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.WithContext(ctx).Model(model).Association(key).Replace(replaceModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLReplaceAssociation, err)
	}

	return nil
}

// RemoveAssociationTx -.
func (r *CRUDDatasourceImpl) RemoveAssociationTx(ctx context.Context, key string, model dto.IRepoEntity, removeModel []dto.IRepoEntity, tx domain.ITransactionEvent) error {
	t, ok := tx.GetTx().(*gorm.DB)
	if !ok {
		return domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.WithContext(ctx).Model(model).Association(key).Delete(removeModel)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeSQLRemoveAssociation, err)
	}

	return nil
}
