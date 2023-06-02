package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/domain/user/repository"
)

// ServiceImpl -.
type ServiceImpl struct {
	UserRepo repository.UserRepository
}

// NewServiceImpl -.
func NewServiceImpl(userRepo repository.UserRepository) *ServiceImpl {
	return &ServiceImpl{UserRepo: userRepo}
}

var ErrUserAlreadyExists = errors.New("user already exists")

func (u *ServiceImpl) RegisterUseCase(ctx context.Context, user *entity.User) (*entity.User, error) {
	// Check if user already exists
	existingUser, err := u.UserRepo.GetByID(ctx, user)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, fmt.Errorf("%w", ErrUserAlreadyExists)
	}

	// Create user
	createdUser, err := u.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

var ErrUserNotFound = errors.New("user not found")

func (u *ServiceImpl) GetByIDUseCase(ctx context.Context, id string) (*entity.User, error) {
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

	return foundUser, nil
}
