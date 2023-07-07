package dto

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/copier"
	"github.com/lib/pq"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

var _ IRepoEntity = (*Role)(nil)

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

func (a *Role) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, NewRoleTransformError(err)
	}

	return a, nil
}

func (a *Role) BackToDomain() (domain.IEntity, error) {
	i := &entity.Role{}
	if err := copier.Copy(i, a); err != nil {
		return nil, NewRoleBackToDomainError(err)
	}

	return i, nil
}

func (a *Role) BeforeUpdate(_ *gorm.DB) (err error) {
	a.UpdatedAt = time.Now()

	return
}

func (a *Role) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID, err = generateID()
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()

	return
}

func (a *Role) GetID() string {
	return a.ID
}

func (a *Role) SetID(id string) {
	a.ID = id
}

func (a *Role) ToJSON() (string, error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", NewRoleToJSONError(err)
	}

	return string(jsonData), nil
}

func (a *Role) DecodeJSON(data string) error {
	err := json.Unmarshal([]byte(data), &a)
	if err != nil {
		return NewRoleDecodeJSONError(err)
	}

	return nil
}
