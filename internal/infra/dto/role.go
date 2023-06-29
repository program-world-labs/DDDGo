package dto

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

var _ entity.IEntity = (*Role)(nil)

type Role struct {
	ID          string         `json:"id" gorm:"type:varchar(20);primary_key"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Permissions pq.StringArray `json:"permissions" gorm:"type:varchar(100)[]"`
	Users       []User         `json:"users" gorm:"many2many:user_roles;"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   time.Time      `json:"deletedAt" gorm:"index"`
}

func (a *Role) TableName() string {
	return "Roles"
}

func (a *Role) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID, err = generateID()

	return
}

func (a *Role) GetID() string {
	return a.ID
}

func (a *Role) SetID(id string) {
	a.ID = id
}

func (a *Role) Self() interface{} {
	return a
}
