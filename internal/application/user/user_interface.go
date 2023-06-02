package usecase

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

type IUserUseCase interface {
	// Command
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
	// Query
	GetByID(ctx context.Context, id string) (*entity.User, error)
}
