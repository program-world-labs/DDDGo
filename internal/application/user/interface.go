package user

import (
	"context"
)

type IService interface {
	// Command
	CreateUser(ctx context.Context, UserInfo *CreatedInput) (*Output, error)
	// AssignRole(ctx context.Context, UserInfo *AssignedInput) (*Output, error)
	UpdateUser(ctx context.Context, UserInfo *UpdatedInput) (*Output, error)
	DeleteUser(ctx context.Context, UserInfo *DeletedInput) (*Output, error)
	// // Query
	GetUserList(ctx context.Context, UserInfo *ListGotInput) (*OutputList, error)
	GetUserDetail(ctx context.Context, UserInfo *DetailGotInput) (*Output, error)
}
