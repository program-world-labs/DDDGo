package user

import (
	"context"

	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

// mockgen -source=internal/application/user/user_interface.go -destination=tests/user/UserService_mock.go -package=mock

type IUserService interface {
	// Command
	RegisterUseCase(ctx context.Context, user *entity.User) (*Output, error)
	// Query
	GetByIDUseCase(ctx context.Context, id string) (*Output, error)
}
