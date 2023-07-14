package repository

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
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
		return nil, domainerrors.Wrap(ErrorCodeRepoTransform, err)
	}

	// 先從Local Cache取得資料
	data, err := r.Cache.Get(ctx, info)
	if err == nil {
		d, derr := data.BackToDomain()
		if derr != nil {
			return nil, domainerrors.Wrap(ErrorCodeRepoBackToDomain, err)
		}

		return d, nil
	}

	// 從Redis取得資料, 取不到資料會自動從db取得資料
	data, err = r.Redis.Get(ctx, info)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	// 將資料存入Local Cache
	_, err = r.Cache.Set(ctx, data)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	d, err := data.BackToDomain()
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepoBackToDomain, err)
	}

	return d, nil
}

// GetAll -.
func (r *CRUDImpl) GetAll(ctx context.Context, sq *domain.SearchQuery, e domain.IEntity) (*domain.List, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepoTransform, err)
	}

	// Define a helper function to handle data

	// Try to get data from local cache
	data, err := r.Cache.GetListItem(ctx, info, sq)
	if err == nil {
		return data.BackToDomain(info)
	}

	// If not in local cache, try to get data from Redis
	data, err = r.Redis.GetListItem(ctx, info, sq)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	return data.BackToDomain(info)
}

// Create -.
func (r *CRUDImpl) Create(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepoTransform, err)
	}

	_, err = r.DB.Create(ctx, info)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepoCreate, err)
	}

	d, err := info.BackToDomain()
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepoBackToDomain, err)
	}

	return d, nil
}

func (r *CRUDImpl) performUpdate(ctx context.Context, e domain.IEntity, action func(ctx context.Context, info dto.IRepoEntity) (dto.IRepoEntity, error)) (domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepoTransform, err)
	}

	data, err := action(ctx, info)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	if data == nil {
		return nil, nil
	}

	d, err := data.BackToDomain()
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepoBackToDomain, err)
	}

	return d, nil
}

// Update -.
func (r *CRUDImpl) Update(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	return r.performUpdate(ctx, e, r.DB.Update)
}

func (r *CRUDImpl) UpdateWithFields(ctx context.Context, e domain.IEntity, keys []string) (domain.IEntity, error) {
	return r.performUpdate(ctx, e, func(ctx context.Context, info dto.IRepoEntity) (dto.IRepoEntity, error) {
		return nil, r.DB.UpdateWithFields(ctx, info, keys)
	})
}

// Delete -.
func (r *CRUDImpl) Delete(ctx context.Context, e domain.IEntity) error {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeRepoTransform, err)
	}

	err = r.DB.Delete(ctx, info)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	return nil
}

// CreateTx -.
func (r *CRUDImpl) CreateTx(ctx context.Context, e domain.IEntity, tx domain.ITransactionEvent) (domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepoTransform, err)
	}

	_, err = r.DB.CreateTx(ctx, info, tx)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepoCreateTx, err)
	}

	d, err := info.BackToDomain()
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepoBackToDomain, err)
	}

	return d, nil
}

func (r *CRUDImpl) UpdateTx(ctx context.Context, e domain.IEntity, tx domain.ITransactionEvent) (domain.IEntity, error) {
	return r.performUpdate(ctx, e, func(ctx context.Context, info dto.IRepoEntity) (dto.IRepoEntity, error) {
		return r.DB.UpdateTx(ctx, info, tx)
	})
}

func (r *CRUDImpl) UpdateWithFieldsTx(ctx context.Context, e domain.IEntity, keys []string, tx domain.ITransactionEvent) error {
	_, err := r.performUpdate(ctx, e, func(ctx context.Context, info dto.IRepoEntity) (dto.IRepoEntity, error) {
		return nil, r.DB.UpdateWithFieldsTx(ctx, info, keys, tx)
	})

	return err
}

// DeleteTx -.
func (r *CRUDImpl) DeleteTx(ctx context.Context, e domain.IEntity, tx domain.ITransactionEvent) error {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeRepoTransform, err)
	}

	err = r.DB.DeleteTx(ctx, info, tx)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeDatasource, err)
	}

	return nil
}
