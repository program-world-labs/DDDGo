package entity

import (
	"time"

	"github.com/program-world-labs/DDDGo/internal/domain"
)

var _ domain.IEntity = (*Group)(nil)

type Group struct {
	ID          string    `json:"id"`          // Group ID
	Name        string    `json:"name"`        // Group Name
	Description string    `json:"description"` // Group Descript
	Users       []User    `json:"users"`       // Group User List
	Owner       *User     `json:"owner"`
	Metadata    string    `json:"metadata"` // json content
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func (a *Group) GetID() string {
	return a.ID
}

func (a *Group) SetID(id string) {
	a.ID = id
}
