// BEGIN: 3c5f9d7d3d9a
package entity

import "github.com/google/uuid"

// User -.
type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	EMail       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	Avatar      string    `json:"avatar"`
}

func NewUser(uid string) (*User, error) {
	uuidString, err := uuid.Parse(uid)
	if err != nil {
		return nil, err
	}

	return &User{
		ID: uuidString,
	}, nil
}

// GetID -.
func (u *User) GetID() string {
	return u.ID.String()
}

// SetID -.
func (u *User) SetID(id string) error {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	u.ID = parsedID

	return nil
}

// TableName -.
func (u *User) TableName() string {
	return "users"
}
