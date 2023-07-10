package repository

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
)

var _ domain.ICRUDRepository = (*CRUDImpl)(nil)

// CRUDImpl -.
type CRUDImpl struct {
	DB        datasource.IDataSource
	Redis     datasource.ICacheDataSource
	Cache     datasource.ICacheDataSource
	DTOEntity dto.IRepoEntity
}

// NewCRUDImpl -.
func NewCRUDImpl(db datasource.IDataSource, redis datasource.ICacheDataSource, cache datasource.ICacheDataSource, data dto.IRepoEntity) *CRUDImpl {
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
		d, derr := data.BackToDomain()
		if derr != nil {
			return nil, NewDatasourceError(err)
		}

		return d, nil
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

	d, err := data.BackToDomain()
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	return d, nil
}

// GetAll -.
func (r *CRUDImpl) GetAll(ctx context.Context, e domain.IEntity, sq *domain.SearchQuery) ([]domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	list, err := r.DB.GetAll(ctx, info, sq)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	var listAll []domain.IEntity

	for _, v := range list {
		d, err := v.BackToDomain()
		if err != nil {
			return nil, NewDatasourceError(err)
		}

		listAll = append(listAll, d)
	}

	return listAll, nil
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

	d, err := info.BackToDomain()
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	return d, nil
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

	d, err := info.BackToDomain()
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	return d, nil
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

	d, err := info.BackToDomain()
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	return d, nil
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

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}

	return nil
}

// CreateTx -.
func (r *CRUDImpl) CreateTx(ctx context.Context, e domain.IEntity, tx domain.ITransactionEvent) (domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	_, err = r.DB.CreateTx(ctx, info, tx)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	d, err := info.BackToDomain()
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	return d, nil
}

// UpdateTx -.
func (r *CRUDImpl) UpdateTx(ctx context.Context, e domain.IEntity, tx domain.ITransactionEvent) (domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	_, err = r.DB.UpdateTx(ctx, info, tx)
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

	d, err := info.BackToDomain()
	if err != nil {
		return nil, NewDatasourceError(err)
	}

	return d, nil
}

// UpdateWithFieldsTx -.
func (r *CRUDImpl) UpdateWithFieldsTx(ctx context.Context, e domain.IEntity, keys []string, tx domain.ITransactionEvent) error {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return NewDatasourceError(err)
	}

	err = r.DB.UpdateWithFieldsTx(ctx, info, keys, tx)
	if err != nil {
		return NewDatasourceError(err)
	}

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}

	return nil
}

// DeleteTx -.
func (r *CRUDImpl) DeleteTx(ctx context.Context, e domain.IEntity, tx domain.ITransactionEvent) error {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return NewDatasourceError(err)
	}

	err = r.DB.DeleteTx(ctx, info, tx)
	if err != nil {
		return NewDatasourceError(err)
	}

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return NewDatasourceError(err)
	}

	return nil
}
