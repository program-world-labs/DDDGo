package currency

import (
	"context"

	"github.com/program-world-labs/pwlogger"
	"go.opentelemetry.io/otel"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/entity"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
	"github.com/program-world-labs/DDDGo/internal/domain/repository"
)

var _ IService = (*ServiceImpl)(nil)

// ServiceImpl -.
type ServiceImpl struct {
	TransactionRepo domain.ITransactionRepo
	CurrencyRepo    repository.CurrencyRepository
	UserRepo        repository.UserRepository
	EventProducer   event.Producer
	log             pwlogger.Interface
}

// NewServiceImpl -.
func NewServiceImpl(currencyRepo repository.CurrencyRepository, userRepo repository.UserRepository, transactionRepo domain.ITransactionRepo, eventProducer event.Producer, l pwlogger.Interface) *ServiceImpl {
	return &ServiceImpl{CurrencyRepo: currencyRepo, UserRepo: userRepo, TransactionRepo: transactionRepo, EventProducer: eventProducer, log: l}
}

// CreateCurrency creates a currency.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) CreateCurrency(ctx context.Context, currencyInfo *CreatedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-createCurrency")

	defer span.End()
	// Validate input.
	err := currencyInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Create currency.
	e := currencyInfo.ToEntity()

	createdCurrency, err := u.CurrencyRepo.Create(ctx, e)

	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Currency.
	createdCurrencyEntity, ok := createdCurrency.(*entity.Currency)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(createdCurrencyEntity), nil
}

// GetCurrencyList gets currency list.
func (u *ServiceImpl) GetCurrencyList(ctx context.Context, currencyInfo *ListGotInput) (*OutputList, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-getCurrencyList")

	defer span.End()
	// Validate input.
	err := currencyInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	sq := currencyInfo.ToSearchQuery()

	// Get currency list.
	list, err := u.CurrencyRepo.GetAll(ctx, sq, &entity.Currency{})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	return NewListOutput(list), nil
}

// GetCurrencyDetail gets currency detail.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) GetCurrencyDetail(ctx context.Context, currencyInfo *DetailGotInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-getCurrencyDetail")

	defer span.End()
	// Validate input.
	err := currencyInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Get currency detail.
	currency, err := u.CurrencyRepo.GetByID(ctx, &entity.Currency{ID: currencyInfo.ID})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Currency.
	currencyEntity, ok := currency.(*entity.Currency)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(currencyEntity), nil
}

// UpdateCurrency updates currency.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) UpdateCurrency(ctx context.Context, currencyInfo *UpdatedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-updateCurrency")

	defer span.End()
	// Validate input.
	err := currencyInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Update currency.
	e := currencyInfo.ToEntity()

	updatedCurrency, err := u.CurrencyRepo.Update(ctx, e)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Currency.
	updatedCurrencyEntity, ok := updatedCurrency.(*entity.Currency)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(updatedCurrencyEntity), nil
}

// DeleteCurrency deletes currency.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) DeleteCurrency(ctx context.Context, currencyInfo *DeletedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-deleteCurrency")

	defer span.End()
	// Validate input.
	err := currencyInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Delete currency.
	info, err := u.CurrencyRepo.Delete(ctx, &entity.Currency{ID: currencyInfo.ID})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Currency.
	currencyEntity, ok := info.(*entity.Currency)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(currencyEntity), nil
}
