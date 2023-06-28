package dto

import (
	"time"

	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

var _ entity.IEntity = (*User)(nil)

type User struct {
	entity.Base
	ID          string    `json:"id" gorm:"primary_key"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	EMail       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	Avatar      string    `json:"avatar"`
	Enabled     bool      `json:"enabled"`
	Roles       []Role    `json:"roles" gorm:"many2many:user_roles;"`
	Wallets     []Wallet  `json:"wallets" gorm:"foreignKey:UserID"`
	Group       Group     `json:"group"`
	GroupID     string    `json:"groupId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt" gorm:"index"`
}

func (a *User) TableName() string {
	return "User"
}

func (a *User) GetID() string {
	return a.ID
}

func (a *User) SetID(id string) {
	a.ID = id
}
