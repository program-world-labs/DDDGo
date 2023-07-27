package group

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
	GroupRepo       repository.GroupRepository
	UserRepo        repository.UserRepository
	EventProducer   event.Producer
	log             pwlogger.Interface
}

// NewServiceImpl -.
func NewServiceImpl(groupRepo repository.GroupRepository, userRepo repository.UserRepository, transactionRepo domain.ITransactionRepo, eventProducer event.Producer, l pwlogger.Interface) *ServiceImpl {
	return &ServiceImpl{GroupRepo: groupRepo, UserRepo: userRepo, TransactionRepo: transactionRepo, EventProducer: eventProducer, log: l}
}

// CreateGroup creates a group.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) CreateGroup(ctx context.Context, groupInfo *CreatedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-createGroup")

	defer span.End()
	// Validate input.
	err := groupInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Create group.
	e := groupInfo.ToEntity()

	createdGroup, err := u.GroupRepo.Create(ctx, e)

	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Group.
	createdGroupEntity, ok := createdGroup.(*entity.Group)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(createdGroupEntity), nil
}

// GetGroupList gets group list.
func (u *ServiceImpl) GetGroupList(ctx context.Context, groupInfo *ListGotInput) (*OutputList, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-getGroupList")

	defer span.End()
	// Validate input.
	err := groupInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	sq := groupInfo.ToSearchQuery()

	// Get group list.
	list, err := u.GroupRepo.GetAll(ctx, sq, &entity.Group{})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	return NewListOutput(list), nil
}

// GetGroupDetail gets group detail.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) GetGroupDetail(ctx context.Context, groupInfo *DetailGotInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-getGroupDetail")

	defer span.End()
	// Validate input.
	err := groupInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Get group detail.
	group, err := u.GroupRepo.GetByID(ctx, &entity.Group{ID: groupInfo.ID})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Group.
	groupEntity, ok := group.(*entity.Group)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(groupEntity), nil
}

// UpdateGroup updates group.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) UpdateGroup(ctx context.Context, groupInfo *UpdatedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-updateGroup")

	defer span.End()
	// Validate input.
	err := groupInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Update group.
	e := groupInfo.ToEntity()

	updatedGroup, err := u.GroupRepo.Update(ctx, e)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Group.
	updatedGroupEntity, ok := updatedGroup.(*entity.Group)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(updatedGroupEntity), nil
}

// DeleteGroup deletes group.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) DeleteGroup(ctx context.Context, groupInfo *DeletedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-deleteGroup")

	defer span.End()
	// Validate input.
	err := groupInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Delete group.
	info, err := u.GroupRepo.Delete(ctx, &entity.Group{ID: groupInfo.ID})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Group.
	groupEntity, ok := info.(*entity.Group)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(groupEntity), nil
}
