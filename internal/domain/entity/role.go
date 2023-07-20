package entity

import (
	"time"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/aggregate"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/event"
)

var _ domain.IEntity = (*Role)(nil)
var _ aggregate.Handler = (*Role)(nil)

type Role struct {
	aggregate.BaseAggregate
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Permissions []string  `json:"permissions"`
	Users       []User    `json:"users"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func NewRole(name, description string, permissions []string) *Role {
	return &Role{
		Name:        name,
		Description: description,
		Permissions: permissions,
	}
}

func NewRoleFromHistory(events []event.DomainEvent) (*Role, error) {
	role := &Role{}
	err := role.LoadFromHistory(events)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (a *Role) GetID() string {
	return a.ID
}

func (a *Role) SetID(id string) {
	a.ID = id
}

// LoadFromHistory -.
func (a *Role) LoadFromHistory(events []event.DomainEvent) error {
	for i := range events {
		err := a.ApplyEvent(&events[i])
		if err != nil {
			return err
		}
	}

	return nil
}

// ApplyEvent -.
func (a *Role) ApplyEvent(domainEvent *event.DomainEvent) error {
	switch domainEvent.Data.(type) {
	case *event.RoleCreatedEvent:
		return a.applyCreated(domainEvent)
	case *event.RoleDescriptionChangedEvent:
		return a.applyDescriptionChanged(domainEvent)
	case **event.RolePermissionUpdatedEvent:
		return a.applyPermisionUpdated(domainEvent)
	default:
		return nil
	}
}

func (a *Role) applyCreated(domainEvent *event.DomainEvent) error {
	eventData, ok := domainEvent.Data.(*event.RoleCreatedEvent)
	if !ok {
		return domainerrors.Wrap(ErrorCodeCastToEvent, ErrCastToEventFailed)
	}

	a.Name = eventData.Name
	a.Description = eventData.Description

	if a.Permissions == nil {
		a.Permissions = []string{}
	}

	a.Permissions = appendUnique(a.Permissions, eventData.Permissions)

	return nil
}

func (a *Role) applyDescriptionChanged(domainEvent *event.DomainEvent) error {
	eventData, ok := domainEvent.Data.(*event.RoleDescriptionChangedEvent)
	if !ok {
		return domainerrors.Wrap(ErrorCodeCastToEvent, ErrCastToEventFailed)
	}

	a.Description = eventData.Description

	return nil
}

func (a *Role) applyPermisionUpdated(domainEvent *event.DomainEvent) error {
	eventData, ok := domainEvent.Data.(*event.RolePermissionUpdatedEvent)
	if !ok {
		return domainerrors.Wrap(ErrorCodeCastToEvent, ErrCastToEventFailed)
	}

	a.Permissions = eventData.Permissions

	return nil
}

func appendUnique(slice []string, elements []string) []string {
	// Create a map where the keys are the strings in the slice.
	// The values don't matter, so we use empty structs as a memory-efficient placeholder.
	unique := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		unique[s] = struct{}{}
	}

	// Append each element in elements to the slice, but only if it's not already in the map.
	for _, e := range elements {
		if _, exists := unique[e]; !exists {
			slice = append(slice, e)
			unique[e] = struct{}{}
		}
	}

	return slice
}
