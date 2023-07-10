package role

import (
	"context"

	"github.com/program-world-labs/pwlogger"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/domain/user/repository"
)

var _ IService = (*ServiceImpl)(nil)

// ServiceImpl -.
type ServiceImpl struct {
	TransactionRepo domain.ITransactionRepo
	RoleRepo        repository.RoleRepository
	log             pwlogger.Interface
}

// NewServiceImpl -.
func NewServiceImpl(roleRepo repository.RoleRepository, transactionRepo domain.ITransactionRepo, l pwlogger.Interface) *ServiceImpl {
	return &ServiceImpl{RoleRepo: roleRepo, TransactionRepo: transactionRepo, log: l}
}

// CreateRole creates a role.
func (u *ServiceImpl) CreateRole(ctx context.Context, roleInfo *CreatedInput) (*Output, error) {
	// Validate input.
	err := roleInfo.Validate()
	if err != nil {
		return nil, NewValidateInputError(err)
	}

	// Create role.
	e := roleInfo.ToEntity()
	createdRole, err := u.RoleRepo.Create(ctx, e)

	if err != nil {
		return nil, NewRepositoryError(err)
	}

	// Cast to entity.Role.
	createdRoleEntity, ok := createdRole.(*entity.Role)
	if !ok {
		return nil, NewCastError(ErrCastToEntityFailed)
	}

	return NewOutput(createdRoleEntity), nil
}
