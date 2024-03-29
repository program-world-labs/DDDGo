package role

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
	RoleRepo        repository.RoleRepository
	UserRepo        repository.UserRepository
	eventStore      event.Store
	EventProducer   event.Producer
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

// CreateRole creates a role.
func (u *ServiceImpl) CreateRole(ctx context.Context, roleInfo *CreatedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-createRole")

	defer span.End()
	// Validate input.
	err := roleInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Create role.
	e := roleInfo.ToEntity()

	createdRole, err := u.RoleRepo.Create(ctx, e)

	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Role.
	createdRoleEntity, ok := createdRole.(*entity.Role)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	// Create domain event
	permissions := make([]string, 0)
	createdEvent := &event.RoleCreatedEvent{
		Name:        roleInfo.Name,
		Description: roleInfo.Description,
		Permissions: append(permissions, createdRoleEntity.Permissions...),
	}
	_, et := createdRoleEntity.GetTypeName(createdRoleEntity)
	domainEvent := event.NewDomainEvent(createdRoleEntity.GetID(), et, -1, createdEvent)

	// Apply event to entity
	err = createdRoleEntity.ApplyEventHelper(createdRoleEntity, domainEvent, true)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeApplyEvent, err, span)
	}

	// Save AccountCreatedEvent to EventStore
	// err = u.eventStore.SafeStore(ctx, createdRoleEntity.Events, createdRoleEntity.Version)
	err = u.eventStore.Store(ctx, createdRoleEntity.Events, 0)
	if err != nil {
		return nil, err
	}

	// Clear Aggregate Uncommit Events
	createdRoleEntity.ClearUnCommitedEvents()

	// Publish event
	err = u.EventProducer.PublishEvent(ctx, createdRoleEntity.Type, domainEvent)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodePublishEvent, err, span)
	}

	return NewOutput(createdRoleEntity), nil
}

// GetRoleList gets role list.
func (u *ServiceImpl) GetRoleList(ctx context.Context, roleInfo *ListGotInput) (*OutputList, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-getRoleList")

	defer span.End()
	// Validate input.
	err := roleInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	sq := roleInfo.ToSearchQuery()

	// Get role list.
	list, err := u.RoleRepo.GetAll(ctx, sq, &entity.Role{})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	return NewListOutput(list), nil
}

// GetRoleDetail gets role detail.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) GetRoleDetail(ctx context.Context, roleInfo *DetailGotInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-getRoleDetail")

	defer span.End()
	// Validate input.
	err := roleInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Get role detail.
	role, err := u.RoleRepo.GetByID(ctx, &entity.Role{ID: roleInfo.ID})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Role.
	roleEntity, ok := role.(*entity.Role)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(roleEntity), nil
}

// UpdateRole updates role.
func (u *ServiceImpl) UpdateRole(ctx context.Context, roleInfo *UpdatedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-updateRole")

	defer span.End()
	// Validate input.
	err := roleInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Update role.
	e := roleInfo.ToEntity()

	updatedRole, err := u.RoleRepo.Update(ctx, e)
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Role.
	updatedRoleEntity, ok := updatedRole.(*entity.Role)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(updatedRoleEntity), nil
}

// DeleteRole deletes role.
//
//nolint:dupl // business logic is different
func (u *ServiceImpl) DeleteRole(ctx context.Context, roleInfo *DeletedInput) (*Output, error) {
	// 開始追蹤
	var tracer = otel.Tracer(domainerrors.GruopID)
	ctx, span := tracer.Start(ctx, "usecase-deleteRole")

	defer span.End()
	// Validate input.
	err := roleInfo.Validate()
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeValidateInput, err, span)
	}

	// Delete role.
	info, err := u.RoleRepo.Delete(ctx, &entity.Role{ID: roleInfo.ID})
	if err != nil {
		return nil, domainerrors.WrapWithSpan(ErrorCodeRepository, err, span)
	}

	// Cast to entity.Role.
	roleEntity, ok := info.(*entity.Role)
	if !ok {
		return nil, domainerrors.WrapWithSpan(ErrorCodeCast, err, span)
	}

	return NewOutput(roleEntity), nil
}
