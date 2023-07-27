package user

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
	EventProducer   event.Producer
	RoleRepo        repository.RoleRepository
	UserRepo        repository.UserRepository
	eventStore      event.Store
	log             pwlogger.Interface
}

// NewServiceImpl -.
func NewServiceImpl(roleRepo repository.RoleRepository, userRepo repository.UserRepository, transactionRepo domain.ITransactionRepo, eventProducer event.Producer, eventStore event.Store, l pwlogger.Interface) *ServiceImpl {
	return &ServiceImpl{
		RoleRepo:        roleRepo,
		UserRepo:        userRepo,
		TransactionRepo: transactionRepo,
		EventProducer:   eventProducer,
		eventStore:      eventStore,
		log:             l,
	}
}

// CreateUser creates a user.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) CreateUser(ctx context.Context, userInfo *CreatedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-createUser")

	defer span.End()
	// Validate input.
	err := userInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Create user.
	e := userInfo.ToEntity()

	createdUser, err := u.UserRepo.Create(ctx, e)

	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.User.
	createdUserEntity, ok := createdUser.(*entity.User)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(createdUserEntity), nil
}

// GetUser gets a user.
func (u *ServiceImpl) GetUserDetail(ctx context.Context, userInfo *DetailGotInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-getUser")

	defer span.End()

	// Validate input.
	err := userInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Get user.
	user, err := u.UserRepo.GetByID(ctx, &entity.User{ID: userInfo.ID})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.User.
	userEntity, ok := user.(*entity.User)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(userEntity), nil
}

// UpdateUser updates a user.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) UpdateUser(ctx context.Context, userInfo *UpdatedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-updateUser")

	defer span.End()

	// Validate input.
	err := userInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Update user.
	e := userInfo.ToEntity()

	// Update user.
	user, err := u.UserRepo.Update(ctx, e)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.User.
	userEntity, ok := user.(*entity.User)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(userEntity), nil
}

// DeleteUser deletes a user.
func (u *ServiceImpl) DeleteUser(ctx context.Context, userInfo *DeletedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-deleteUser")

	defer span.End()

	// Validate input.
	err := userInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Delete user.
	info, err := u.UserRepo.Delete(ctx, &entity.User{ID: userInfo.ID})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.User.
	userEntity, ok := info.(*entity.User)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(userEntity), nil
}

// GetUserList lists users.
func (u *ServiceImpl) GetUserList(ctx context.Context, userInfo *ListGotInput) (*OutputList, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-listUsers")

	defer span.End()

	// Validate input.
	err := userInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	sq := userInfo.ToSearchQuery()

	// Get role list.
	list, err := u.UserRepo.GetAll(ctx, sq, &entity.Role{})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	return NewListOutput(list), nil
}
