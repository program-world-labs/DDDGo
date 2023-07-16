package repository

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

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
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "repo-getByID")

	defer span.End()

	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoTransform, err, span)
	}

	// 先從Local Cache取得資料
	data, err := r.Cache.Get(ctx, info)
	if err == nil {
		d, derr := data.BackToDomain()
		if derr != nil {
			return nil, domainerrors.WrapWithSpan(ErrorCodeRepoBackToDomain, err, span)
		}

		return d, nil
	}

	// 從Redis取得資料, 取不到資料會自動從db取得資料
	data, err = r.Redis.Get(ctx, info)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	// 將資料存入Local Cache
	_, err = r.Cache.Set(ctx, data)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	d, err := data.BackToDomain()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoBackToDomain, err, span)
	}

	return d, nil
}

// GetAll -.
func (r *CRUDImpl) GetAll(ctx context.Context, sq *domain.SearchQuery, e domain.IEntity) (*domain.List, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "repo-getAll")

	defer span.End()

	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoTransform, err, span)
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
		return nil, domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	return data.BackToDomain(info)
}

// Create -.
func (r *CRUDImpl) Create(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "repo-create")

	defer span.End()

	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoTransform, err, span)
	}

	_, err = r.DB.Create(ctx, info)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoCreate, err, span)
	}

	d, err := info.BackToDomain()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoBackToDomain, err, span)
	}

	return d, nil
}

func (r *CRUDImpl) performUpdate(ctx context.Context, e domain.IEntity, span trace.Span, action func(ctx context.Context, info dto.IRepoEntity) (dto.IRepoEntity, error)) (domain.IEntity, error) {
	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoTransform, err, span)
	}

	data, err := action(ctx, info)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	if data == nil {
		return nil, nil
	}

	d, err := data.BackToDomain()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoBackToDomain, err, span)
	}

	return d, nil
}

// Update -.
func (r *CRUDImpl) Update(ctx context.Context, e domain.IEntity) (domain.IEntity, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "repo-update")

	defer span.End()

	return r.performUpdate(ctx, e, span, r.DB.Update)
}

func (r *CRUDImpl) UpdateWithFields(ctx context.Context, e domain.IEntity, keys []string) (domain.IEntity, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "repo-updateWithFields")

	defer span.End()

	return r.performUpdate(ctx, e, span, func(ctx context.Context, info dto.IRepoEntity) (dto.IRepoEntity, error) {
		return nil, r.DB.UpdateWithFields(ctx, info, keys)
	})
}

// Delete -.
func (r *CRUDImpl) Delete(ctx context.Context, e domain.IEntity) error {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "repo-delete")

	defer span.End()

	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeRepoTransform, err, span)
	}

	err = r.DB.Delete(ctx, info)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	return nil
}

// CreateTx -.
func (r *CRUDImpl) CreateTx(ctx context.Context, e domain.IEntity, tx domain.ITransactionEvent) (domain.IEntity, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "repo-createTx")

	defer span.End()

	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoTransform, err, span)
	}

	_, err = r.DB.CreateTx(ctx, info, tx)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoCreateTx, err, span)
	}

	d, err := info.BackToDomain()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepoBackToDomain, err, span)
	}

	return d, nil
}

func (r *CRUDImpl) UpdateTx(ctx context.Context, e domain.IEntity, tx domain.ITransactionEvent) (domain.IEntity, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "repo-updateTx")

	defer span.End()

	return r.performUpdate(ctx, e, span, func(ctx context.Context, info dto.IRepoEntity) (dto.IRepoEntity, error) {
		return r.DB.UpdateTx(ctx, info, tx)
	})
}

func (r *CRUDImpl) UpdateWithFieldsTx(ctx context.Context, e domain.IEntity, keys []string, tx domain.ITransactionEvent) error {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "repo-updateWithFieldsTx")

	defer span.End()

	_, err := r.performUpdate(ctx, e, span, func(ctx context.Context, info dto.IRepoEntity) (dto.IRepoEntity, error) {
		return nil, r.DB.UpdateWithFieldsTx(ctx, info, keys, tx)
	})

	return err
}

// DeleteTx -.
func (r *CRUDImpl) DeleteTx(ctx context.Context, e domain.IEntity, tx domain.ITransactionEvent) error {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "repo-deleteTx")

	defer span.End()

	info, err := r.DTOEntity.Transform(e)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeRepoTransform, err, span)
	}

	err = r.DB.DeleteTx(ctx, info, tx)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	err = r.Cache.Delete(ctx, info)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	err = r.Redis.Delete(ctx, info)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeDatasource, err, span)
	}

	return nil
}
