package entity

import (
	"errors"
	"time"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/aggregate"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
)

var _ domain.IEntity = (*User)(nil)
var _ aggregate.AggregateHandler = (*User)(nil)

// User -.
type User struct {
	aggregate.BaseAggregate
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	EMail       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	Avatar      string    `json:"avatar"`
	Roles       []Role    `json:"roles" gorm:"many2many:user_roles;"`
	Department  Group     `json:"departmentId" gorm:"foreignKey:DepartmentID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

func NewUser(uid string) (*User, error) {
	return &User{
		ID: uid,
	}, nil
}

func NewUserFromHistory(events []event.DomainEvent) (*User, error) {
	user := &User{}
	err := user.LoadFromHistory(events)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetID -.
func (u *User) GetID() string {
	return u.ID
}

// SetID -.
func (u *User) SetID(id string) {
	u.ID = id
}

// LoadFromHistory -.
func (u *User) LoadFromHistory(events []event.DomainEvent) error {
	for i := range events {
		err := u.ApplyEvent(&events[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// ApplyEvent -.
func (u *User) ApplyEvent(domainEvent *event.DomainEvent) error {
	switch domainEvent.Data.(type) {
	case *event.UserCreatedEvent:
		return u.applyCreated(domainEvent)
	case *event.UserPasswordChangedEvent:
		return u.applyPasswordChanged(domainEvent)
	case *event.UserEmailChangedEvent:
		return u.applyEmailChanged(domainEvent)
	default:
		return nil
	}
}

func (u *User) applyCreated(domainEvent *event.DomainEvent) error {
	eventData, ok := domainEvent.Data.(*event.UserCreatedEvent)
	if !ok {
		return errors.New("invalid event data")
	}

	u.Username = eventData.UserName
	u.Password = eventData.Password
	u.EMail = eventData.EMail

	return nil
}

func (u *User) applyPasswordChanged(domainEvent *event.DomainEvent) error {
	eventData, ok := domainEvent.Data.(*event.UserPasswordChangedEvent)
	if !ok {
		return errors.New("invalid event data")
	}

	u.Password = eventData.Password

	return nil
}

func (u *User) applyEmailChanged(domainEvent *event.DomainEvent) error {
	eventData, ok := domainEvent.Data.(*event.UserEmailChangedEvent)
	if !ok {
		return errors.New("invalid event data")
	}

	u.EMail = eventData.EMail

	return nil
}
