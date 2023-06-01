package usecase

import (
	"context"

	"gitlab.com/demojira/template.git/internal/domain/user/entity"
)

type UserUseCase interface {
	// Command
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
	// Query
	GetByID(ctx context.Context, id string) (*entity.User, error)
}
