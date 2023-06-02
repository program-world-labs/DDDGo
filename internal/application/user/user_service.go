package user

import (
	"context"
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/domain/user/repository"
)

// UserServiceImpl -.
type UserServiceImpl struct {
	UserRepo repository.UserRepository
}

// NewUserServiceImpl -.
func NewUserServiceImpl(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepo: userRepo}
}

// Register -.
func (u *UserServiceImpl) RegisterUseCase(ctx context.Context, user *entity.User) (*entity.User, error) {
	// Check if user already exists
	existingUser, err := u.UserRepo.GetByID(ctx, user)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Create user
	createdUser, err := u.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// GetByID -.
func (u *UserServiceImpl) GetByIDUseCase(ctx context.Context, id string) (*entity.User, error) {
	user, err := entity.NewUser(id)
	if err != nil {
		return nil, err
	}
	foundUser, err := u.UserRepo.GetByID(ctx, user)
	if err != nil {
		return nil, err
	}
	if foundUser == nil {
		return nil, errors.New("user not found")
	}
	return foundUser, nil
}
