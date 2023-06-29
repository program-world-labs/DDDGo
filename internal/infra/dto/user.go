package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

var _ entity.IEntity = (*User)(nil)

type User struct {
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
	return "Users"
}

func (a *User) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID = uuid.New().String()

	return
}

func (a *User) GetID() string {
	return a.ID
}

func (a *User) SetID(id string) {
	a.ID = id
}

func (a *User) Self() interface{} {
	return a
}
