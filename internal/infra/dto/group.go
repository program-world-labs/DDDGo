package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

var _ entity.IEntity = (*Group)(nil)

type Group struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Users       []User    `json:"users" gorm:"foreignKey:GroupID"`
	Owner       *User     `json:"owner"`
	Metadata    string    `json:"metadata"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt" gorm:"index"`
}

func (a *Group) TableName() string {
	return "Groups"
}

func (a *Group) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID = uuid.New().String()

	return
}

func (a *Group) GetID() string {
	return a.ID
}

func (a *Group) SetID(id string) {
	a.ID = id
}

func (a *Group) Self() interface{} {
	return a
}
