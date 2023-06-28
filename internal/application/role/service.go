package role

import (
	"context"
	"fmt"

	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/domain/user/repository"
	"github.com/program-world-labs/pwlogger"
)

var _ IService = (*ServiceImpl)(nil)

// ServiceImpl -.
type ServiceImpl struct {
	RoleRepo repository.RoleRepository
	log      pwlogger.Interface
}

// NewServiceImpl -.
func NewServiceImpl(roleRepo repository.RoleRepository, l pwlogger.Interface) *ServiceImpl {
	return &ServiceImpl{RoleRepo: roleRepo, log: l}
}

// CreateRole -.
func (u *ServiceImpl) CreateRole(ctx context.Context, roleInfo *CreatedInput) (*Output, error) {
	e := roleInfo.ToEntity()
	// Create role
	createdRole, err := u.RoleRepo.Create(ctx, e)
	if err != nil {
		return nil, err
	}

	// Cast to entity.Role
	createdRoleEntity, ok := createdRole.(*entity.Role)
	if !ok {
		return nil, fmt.Errorf("ServiceImpl - CreateRole - u.RoleRepo.Create: %w", err)
	}

	return NewOutput(createdRoleEntity), nil
}
