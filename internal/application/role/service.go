package role

import (
	"context"

	"github.com/program-world-labs/pwlogger"

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
	EventProducer   event.EventProducer
	log             pwlogger.Interface
}

// NewServiceImpl -.
func NewServiceImpl(roleRepo repository.RoleRepository, transactionRepo domain.ITransactionRepo, eventProducer event.EventProducer, l pwlogger.Interface) *ServiceImpl {
	return &ServiceImpl{RoleRepo: roleRepo, TransactionRepo: transactionRepo, EventProducer: eventProducer, log: l}
}

// CreateRole creates a role.
func (u *ServiceImpl) CreateRole(ctx context.Context, roleInfo *CreatedInput) (*Output, error) {
	// Validate input.
	err := roleInfo.Validate()
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeValidateInput, err)
	}

	// Create role.
	e := roleInfo.ToEntity()

	createdRole, err := u.RoleRepo.Create(ctx, e)

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepository, err)
	}

	// Cast to entity.Role.
	createdRoleEntity, ok := createdRole.(*entity.Role)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeCast, err)
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
	createdRoleEntity.ApplyEventHelper(createdRoleEntity, domainEvent, true)

	// Publish event
	err = u.EventProducer.PublishEvent(createdRoleEntity.Type, domainEvent)
	if err != nil {
		return nil, err
	}

	return NewOutput(createdRoleEntity), nil
}

// GetRoleList gets role list.
func (u *ServiceImpl) GetRoleList(ctx context.Context, roleInfo *ListGotInput) (*OutputList, error) {
	// Validate input.
	err := roleInfo.Validate()
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeValidateInput, err)
	}

	sq := roleInfo.ToSearchQuery()

	// Get role list.
	list, err := u.RoleRepo.GetAll(ctx, sq, &entity.Role{})
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepository, err)
	}

	return NewListOutput(list), nil
}

// GetRoleDetail gets role detail.
func (u *ServiceImpl) GetRoleDetail(ctx context.Context, roleInfo *DetailGotInput) (*Output, error) {
	// Validate input.
	err := roleInfo.Validate()
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeValidateInput, err)
	}

	// Get role detail.
	role, err := u.RoleRepo.GetByID(ctx, &entity.Role{ID: roleInfo.ID})
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRepository, err)
	}

	// Cast to entity.Role.
	roleEntity, ok := role.(*entity.Role)
	if !ok {
		return nil, domainerrors.Wrap(ErrorCodeCast, err)
	}

	return NewOutput(roleEntity), nil
}
