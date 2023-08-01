package nosqlfs

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/infra/datasource"
	"github.com/program-world-labs/DDDGo/internal/infra/dto"
	firestoredb "github.com/program-world-labs/DDDGo/pkg/pwsql/nosql/firestoreDB"
)

var _ datasource.IDataSource = (*CRUDDatasourceImpl)(nil)

// CRUDDatasourceImpl -.
type CRUDDatasourceImpl struct {
	Client *firestore.Client
}

// NewCRUDDatasourceImpl -.
func NewCRUDDatasourceImpl(db *firestoredb.Firestore) *CRUDDatasourceImpl {
	return &CRUDDatasourceImpl{Client: db.GetClient()}
}

// GetByID -.
func (r *CRUDDatasourceImpl) GetByID(ctx context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	doc, err := r.Client.Collection(model.TableName()).Doc(model.GetID()).Get(ctx)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLGet, err)
	}

	err = doc.DataTo(model)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLGet, err)
	}

	return model, nil
}

// Create -.
func (r *CRUDDatasourceImpl) Create(ctx context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	// BeforeCreate
	if err := model.BeforeCreate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLCreate, err)
	}

	_, err := r.Client.Collection(model.TableName()).Doc(model.GetID()).Create(ctx, model)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLCreate, err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl) Update(ctx context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	// BeforeUpdate
	if err := model.BeforeUpdate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	_, err := r.Client.Collection(model.TableName()).Doc(model.GetID()).Set(ctx, model)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	return model, nil
}

// UpdateWithFields -.
func (r *CRUDDatasourceImpl) UpdateWithFields(ctx context.Context, model dto.IRepoEntity, fields []string) (dto.IRepoEntity, error) {
	// BeforeUpdate
	if err := model.BeforeUpdate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdateWithFields, err)
	}

	_, err := r.Client.Collection(model.TableName()).Doc(model.GetID()).Set(ctx, model, firestore.Merge(fields))
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdateWithFields, err)
	}

	return model, nil
}

// Delete -.
func (r *CRUDDatasourceImpl) Delete(ctx context.Context, model dto.IRepoEntity) (dto.IRepoEntity, error) {
	_, err := r.Client.Collection(model.TableName()).Doc(model.GetID()).Delete(ctx)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLDelete, err)
	}

	return model, nil
}

func (r *CRUDDatasourceImpl) GetAll(ctx context.Context, sq *domain.SearchQuery, model dto.IRepoEntity) (*dto.List, error) {
	query := r.Client.Collection(model.TableName()).Query
	// 加入Filter Query
	for _, filter := range sq.Filters {
		query = query.Where(filter.FilterField, filter.Operator, filter.Value)
	}
	// 加入Order Query
	for _, order := range sq.Orders {
		dir := firestore.Asc
		if order.Dir == "desc" {
			dir = firestore.Desc
		}

		query = query.OrderBy(order.OrderField, dir)
	}
	// 加入Limit Query
	if sq.Page.Limit > 0 {
		query = query.Limit(sq.Page.Limit)
	}
	// 加入Offset Query
	if sq.Page.Offset > 0 {
		query = query.Offset(sq.Page.Offset)
	}

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLGetAll, err)
	}

	list := &dto.List{}
	list.Limit = sq.Page.Limit
	list.Offset = sq.Page.Offset

	dataList := []interface{}{}

	for _, doc := range docs {
		err = doc.DataTo(model)
		if err != nil {
			return nil, domainerrors.Wrap(ErrorCodeSQLGetAll, err)
		}

		dataList = append(dataList, model)
	}

	list.Data = dataList

	if len(docs) == 0 {
		list.Data = []interface{}{}
	}

	return list, nil
}

// Create -.
func (r *CRUDDatasourceImpl) CreateTx(ctx context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	t, ok := tx.GetTx().(*firestore.Transaction)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	// BeforeCreate
	if err := model.BeforeCreate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLCreate, err)
	}

	err := t.Create(r.Client.Collection(model.TableName()).Doc(model.GetID()), model)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLCreate, err)
	}

	return model, nil
}

// Update -.
func (r *CRUDDatasourceImpl) UpdateTx(ctx context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	t, ok := tx.GetTx().(*firestore.Transaction)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	// BeforeUpdate
	if err := model.BeforeUpdate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	err := t.Set(r.Client.Collection(model.TableName()).Doc(model.GetID()), model)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdate, err)
	}

	return model, nil
}

// UpdateWithFields -.
func (r *CRUDDatasourceImpl) UpdateWithFieldsTx(ctx context.Context, model dto.IRepoEntity, fields []string, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	t, ok := tx.GetTx().(*firestore.Transaction)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	// BeforeUpdate
	if err := model.BeforeUpdate(); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdateWithFields, err)
	}

	err := t.Set(r.Client.Collection(model.TableName()).Doc(model.GetID()), model, firestore.Merge(fields))
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLUpdateWithFields, err)
	}

	return model, nil
}

// Delete -.
func (r *CRUDDatasourceImpl) DeleteTx(ctx context.Context, model dto.IRepoEntity, tx domain.ITransactionEvent) (dto.IRepoEntity, error) {
	t, ok := tx.GetTx().(*firestore.Transaction)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeSQLCast, ErrCastToEntityFailed)
	}

	err := t.Delete(r.Client.Collection(model.TableName()).Doc(model.GetID()))
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeSQLDelete, err)
	}

	return model, nil
}
