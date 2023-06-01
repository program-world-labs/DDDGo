package usecase

import (
	"context"
	"errors"

	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
	"github.com/program-world-labs/DDDGo/internal/domain/user/repository"
)

// UserUseCaseImpl -.
type UserUseCaseImpl struct {
	UserRepo repository.UserRepository
}

// NewUserUseCaseImpl -.
func NewUserUseCaseImpl(userRepo repository.UserRepository) *UserUseCaseImpl {
	return &UserUseCaseImpl{UserRepo: userRepo}
}

// Register -.
func (u *UserUseCaseImpl) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
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
func (u *UserUseCaseImpl) GetByID(ctx context.Context, id string) (*entity.User, error) {
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
