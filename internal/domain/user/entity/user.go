package entity

import (
	"time"

	"github.com/program-world-labs/DDDGo/internal/domain"
)

var _ domain.IEntity = (*User)(nil)

// User -.
type User struct {
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

// GetID -.
func (u *User) GetID() string {
	return u.ID
}

// SetID -.
func (u *User) SetID(id string) {
	u.ID = id
}

// Self -.
func (u *User) Self() interface{} {
	return u
}
