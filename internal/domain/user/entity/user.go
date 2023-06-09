// BEGIN: 3c5f9d7d3d9a
package entity

import (
	"github.com/program-world-labs/DDDGo/internal/domain"
)

var _ domain.IEntity = (*User)(nil)

// User -.
type User struct {
	ID          string `gorm:"primary_key;"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	EMail       string `json:"email"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
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
func (u *User) SetID(id string) error {
	u.ID = id

	return nil
}

// TableName -.
func (u *User) TableName() string {
	return "users"
}
