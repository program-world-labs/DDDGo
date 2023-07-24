package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/allegro/bigcache/v3"
	"go.opentelemetry.io/otel"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
)

var _ datasource.ICacheDataSource = (*BigCacheDataSourceImpl)(nil)

// BigCacheDataSourceImpl -.
type BigCacheDataSourceImpl struct {
	Cache *bigcache.BigCache
}

// NewBigCacheDataSourceImpl -.
func NewBigCacheDataSourceImp(cache *bigcache.BigCache) *BigCacheDataSourceImpl {
	return &BigCacheDataSourceImpl{Cache: cache}
}

func (r *BigCacheDataSourceImpl) cacheKey(model dto.IRepoEntity, sq ...*domain.SearchQuery) string {
	if len(sq) > 0 {
		return fmt.Sprintf("%s-%s-%s", model.TableName(), model.GetID(), sq[0].GetKey())
	}

	return fmt.Sprintf("%s-%s", model.TableName(), model.GetID())
}

// Get -.
func (r *BigCacheDataSourceImpl) Get(ctx context.Context, model dto.IRepoEntity, _ ...time.Duration) (dto.IRepoEntity, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	_, span := tracer.Start(ctx, "datasource-get-local")

	defer span.End()

	data, err := r.Cache.Get(r.cacheKey(model))
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCacheGet, err, span)
	}

	err = json.Unmarshal(data, &model)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCacheGet, err, span)
	}

	return model, nil
}

// Set -.
func (r *BigCacheDataSourceImpl) Set(ctx context.Context, model dto.IRepoEntity, _ ...time.Duration) (dto.IRepoEntity, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	_, span := tracer.Start(ctx, "datasource-set-local")

	defer span.End()

	data, err := json.Marshal(model)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCacheSet, err, span)
	}

	err = r.Cache.Set(r.cacheKey(model), data)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCacheSet, err, span)
	}

	return model, nil
}

// Delete -.
func (r *BigCacheDataSourceImpl) Delete(ctx context.Context, model dto.IRepoEntity) error {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	_, span := tracer.Start(ctx, "datasource-delete-local")

	defer span.End()

	err := r.Cache.Delete(r.cacheKey(model))
	if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return domainerrors.WrapWithSpan(ErrorCodeCacheDelete, err, span)
	}

	return nil
}

// GetListItem -.
func (r *BigCacheDataSourceImpl) GetListItem(ctx context.Context, model dto.IRepoEntity, sq *domain.SearchQuery, _ ...time.Duration) (*dto.List, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	_, span := tracer.Start(ctx, "datasource-getListItem-local")

	defer span.End()

	data, err := r.Cache.Get(r.cacheKey(model, sq))
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCacheGet, err, span)
	}

	var result *dto.List
	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCacheGet, err, span)
	}

	// 取得分頁資訊
	page, err := r.Cache.Get(fmt.Sprintf("%s-ListKeys", model.TableName()))
	if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCacheGet, err, span)
	}

	// 加入新的key，儲存所有的List key
	v := fmt.Sprintf("%s,%s", page, r.cacheKey(model, sq))

	err = r.Cache.Set(fmt.Sprintf("%s-ListKeys", model.TableName()), []byte(v))
	if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil, domainerrors.Wrap(ErrorCodeCacheGet, err)
	}

	return result, nil
}

// SetListItem -.
func (r *BigCacheDataSourceImpl) SetListItem(ctx context.Context, model []dto.IRepoEntity, sq *domain.SearchQuery, count int64, _ ...time.Duration) error {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	_, span := tracer.Start(ctx, "datasource-setListItem-local")

	defer span.End()

	data, err := json.Marshal(model)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeCacheSet, err, span)
	}

	var info []map[string]interface{}
	err = json.Unmarshal(data, &info)

	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeCacheSet, err, span)
	}

	var domainList = map[string]interface{}{
		"data":   info,
		"total":  count,
		"limit":  sq.Page.Limit,
		"offset": sq.Page.Offset,
	}

	result, err := json.Marshal(domainList)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeCacheSet, err, span)
	}

	err = r.Cache.Set(r.cacheKey(model[0], sq), result)
	if err != nil {
		return domainerrors.WrapWithSpan(ErrorCodeCacheSet, err, span)
	}

	return nil
}

// GetListKeys -.
func (r *BigCacheDataSourceImpl) GetListKeys(ctx context.Context, model dto.IRepoEntity) ([]string, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	_, span := tracer.Start(ctx, "datasource-getListKeys-local")

	defer span.End()

	data, err := r.Cache.Get(fmt.Sprintf("%s-ListKeys", model.TableName()))
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCacheGet, err, span)
	}

	return strings.Split(string(data), ","), nil
}

// DeleteListKeys -.
func (r *BigCacheDataSourceImpl) DeleteListKeys(ctx context.Context, model dto.IRepoEntity) error {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	_, span := tracer.Start(ctx, "datasource-getListKeys-local")

	defer span.End()

	err := r.Cache.Delete(fmt.Sprintf("%s-ListKeys", model.TableName()))
	if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return domainerrors.WrapWithSpan(ErrorCodeCacheDelete, err, span)
	}

	return nil
}

func (r *BigCacheDataSourceImpl) DeleteWithKey(ctx context.Context, key string) error {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	_, span := tracer.Start(ctx, "datasource-getListKeys-local")

	defer span.End()

	err := r.Cache.Delete(key)
	if err != nil && !errors.Is(err, bigcache.ErrEntryNotFound) {
		return domainerrors.WrapWithSpan(ErrorCodeCacheDelete, err, span)
	}

	return nil
}
