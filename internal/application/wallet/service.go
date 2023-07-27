package wallet

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
	WalletRepo      repository.WalletRepository
	UserRepo        repository.UserRepository
	EventProducer   event.Producer
	log             pwlogger.Interface
}

// NewServiceImpl -.
func NewServiceImpl(walletRepo repository.WalletRepository, userRepo repository.UserRepository, transactionRepo domain.ITransactionRepo, eventProducer event.Producer, l pwlogger.Interface) *ServiceImpl {
	return &ServiceImpl{WalletRepo: walletRepo, UserRepo: userRepo, TransactionRepo: transactionRepo, EventProducer: eventProducer, log: l}
}

// CreateWallet creates a wallet.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) CreateWallet(ctx context.Context, walletInfo *CreatedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-createWallet")

	defer span.End()
	// Validate input.
	err := walletInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Create wallet.
	e := walletInfo.ToEntity()

	createdWallet, err := u.WalletRepo.Create(ctx, e)

	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Wallet.
	createdWalletEntity, ok := createdWallet.(*entity.Wallet)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(createdWalletEntity), nil
}

// GetWalletList gets wallet list.
func (u *ServiceImpl) GetWalletList(ctx context.Context, walletInfo *ListGotInput) (*OutputList, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-getWalletList")

	defer span.End()
	// Validate input.
	err := walletInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	sq := walletInfo.ToSearchQuery()

	// Get wallet list.
	list, err := u.WalletRepo.GetAll(ctx, sq, &entity.Wallet{})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	return NewListOutput(list), nil
}

// GetWalletDetail gets wallet detail.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) GetWalletDetail(ctx context.Context, walletInfo *DetailGotInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-getWalletDetail")

	defer span.End()
	// Validate input.
	err := walletInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Get wallet detail.
	wallet, err := u.WalletRepo.GetByID(ctx, &entity.Wallet{ID: walletInfo.ID})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Wallet.
	walletEntity, ok := wallet.(*entity.Wallet)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(walletEntity), nil
}

// UpdateWallet updates wallet.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) UpdateWallet(ctx context.Context, walletInfo *UpdatedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-updateWallet")

	defer span.End()
	// Validate input.
	err := walletInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Update wallet.
	e := walletInfo.ToEntity()

	updatedWallet, err := u.WalletRepo.Update(ctx, e)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Wallet.
	updatedWalletEntity, ok := updatedWallet.(*entity.Wallet)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(updatedWalletEntity), nil
}

// DeleteWallet deletes wallet.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) DeleteWallet(ctx context.Context, walletInfo *DeletedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-deleteWallet")

	defer span.End()
	// Validate input.
	err := walletInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Delete wallet.
	info, err := u.WalletRepo.Delete(ctx, &entity.Wallet{ID: walletInfo.ID})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Wallet.
	walletEntity, ok := info.(*entity.Wallet)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(walletEntity), nil
}
