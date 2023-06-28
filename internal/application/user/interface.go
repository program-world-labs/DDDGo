package user

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

type IService interface {
	// Command
	RegisterUseCase(ctx context.Context, userInfo *entity.User) (*Output, error)
	// Query
	GetByIDUseCase(ctx context.Context, id string) (*Output, error)
}
