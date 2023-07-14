package role

import (
	"context"
)

type IService interface {
	// Command
	CreateRole(ctx context.Context, roleInfo *CreatedInput) (*Output, error)
	// AssignRole(ctx context.Context, roleInfo *AssignedInput) (*Output, error)
	UpdateRole(ctx context.Context, roleInfo *UpdatedInput) (*Output, error)
	// DeleteRole(ctx context.Context, id string) error
	// // Query
	GetRoleList(ctx context.Context, roleInfo *ListGotInput) (*OutputList, error)
	GetRoleDetail(ctx context.Context, roleInfo *DetailGotInput) (*Output, error)
}
