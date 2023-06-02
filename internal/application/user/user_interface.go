package user

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

type IUserService interface {
	// Command
	RegisterUseCase(ctx context.Context, user *entity.User) (*entity.User, error)
	// Query
	GetByIDUseCase(ctx context.Context, id string) (*entity.User, error)
}
