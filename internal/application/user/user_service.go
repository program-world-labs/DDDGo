package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/domain/user/repository"
	"github.com/program-world-labs/DDDGo/pkg/logger"
	"github.com/program-world-labs/DDDGo/pkg/operations"
)

var _ IUserService = (*ServiceImpl)(nil)

// ServiceImpl -.
type ServiceImpl struct {
	UserRepo repository.UserRepository
	log      logger.Interface
	trace    operations.ITracer
}

// NewServiceImpl -.
func NewServiceImpl(userRepo repository.UserRepository, l logger.Interface, t operations.ITracer) *ServiceImpl {
	return &ServiceImpl{UserRepo: userRepo, log: l, trace: t}
}

var ErrUserAlreadyExists = errors.New("user already exists")

func (u *ServiceImpl) RegisterUseCase(ctx context.Context, userInfo *entity.User) (*Output, error) {
	// Check if user already exists
	existingUser, err := u.UserRepo.GetByID(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, fmt.Errorf("%w", ErrUserAlreadyExists)
	}

	// Create user
	createdUser, err := u.UserRepo.Create(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	// Cast to entity.User
	createdUserEntity, ok := createdUser.(*entity.User)
	if !ok {
		return nil, fmt.Errorf("ServiceImpl - RegisterUseCase - u.UserRepo.Create: %w", err)
	}

	return NewOutput(createdUserEntity), nil
}

var ErrUserNotFound = errors.New("user not found")

func (u *ServiceImpl) GetByIDUseCase(ctx context.Context, id string) (*Output, error) {
	user, err := entity.NewUser(id)
	if err != nil {
		return nil, err
	}

	foundUser, err := u.UserRepo.GetByID(ctx, user)
	if err != nil {
		return nil, err
	}

	if foundUser == nil {
		return nil, fmt.Errorf("%w", ErrUserNotFound)
	}

	// cast to entity.User
	foundUserEntity, ok := foundUser.(*entity.User)
	if !ok {
		return nil, fmt.Errorf("ServiceImpl - GetByIDUseCase - u.UserRepo.GetByID: %w", err)
	}

	return NewOutput(foundUserEntity), nil
}
