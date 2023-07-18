package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/program-world-labs/pwlogger"

	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/entity"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
	"github.com/program-world-labs/DDDGo/internal/domain/repository"
)

var _ IService = (*ServiceImpl)(nil)

// ServiceImpl -.
type ServiceImpl struct {
	UserRepo      repository.UserRepository
	EventProducer event.EventProducer
	log           pwlogger.Interface
}

// NewServiceImpl -.
func NewServiceImpl(userRepo repository.UserRepository, eventProducer event.EventProducer, l pwlogger.Interface) *ServiceImpl {
	return &ServiceImpl{UserRepo: userRepo, EventProducer: eventProducer, log: l}
}

var ErrUserAlreadyExists = errors.New("user already exists")

func (u *ServiceImpl) RegisterUseCase(ctx context.Context, userInfo *entity.User) (*Output, error) {
	// Check Input Format
	if err := validator.New().Struct(userInfo); err != nil {
		return nil, domainerrors.New(fmt.Sprint(ErrorCodeValidateInput), err.Error())
	}

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

	// Create domain event
	createdEvent := &event.UserCreatedEvent{
		UserName: userInfo.Username,
		Password: userInfo.Password,
		EMail:    userInfo.EMail,
	}
	_, et := createdUserEntity.GetTypeName(createdEvent)
	e := event.NewDomainEvent(createdUserEntity.GetID(), et, -1, createdEvent)

	// Apply event to entity
	createdUserEntity.ApplyEventHelper(createdUserEntity, e, true)

	// Publish event
	err = u.EventProducer.PublishEvent(createdUserEntity.Type, e)
	if err != nil {
		return nil, err
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
