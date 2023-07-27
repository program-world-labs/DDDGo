package group

import (
	"context"
)

type IService interface {
	// Command
	CreateGroup(ctx context.Context, groupInfo *CreatedInput) (*Output, error)
	// AssignGroup(ctx context.Context, groupInfo *AssignedInput) (*Output, error)
	UpdateGroup(ctx context.Context, groupInfo *UpdatedInput) (*Output, error)
	DeleteGroup(ctx context.Context, groupInfo *DeletedInput) (*Output, error)
	// // Query
	GetGroupList(ctx context.Context, groupInfo *ListGotInput) (*OutputList, error)
	GetGroupDetail(ctx context.Context, groupInfo *DetailGotInput) (*Output, error)
}
