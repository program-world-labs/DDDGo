package dto

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

var _ IRepoEntity = (*Group)(nil)

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

func (a *Group) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, NewGroupTransformError(err)
	}

	return a, nil
}

func (a *Group) BackToDomain() (domain.IEntity, error) {
	i := &entity.Group{}
	if err := copier.Copy(&i, a); err != nil {
		return nil, NewGroupBackToDomainError(err)
	}

	return i, nil
}

func (a *Group) BeforeUpdate(_ *gorm.DB) (err error) {
	a.UpdatedAt = time.Now()

	return
}
func (a *Group) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID, err = generateID()
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()

	return
}

func (a *Group) GetID() string {
	return a.ID
}

func (a *Group) SetID(id string) {
	a.ID = id
}

func (a *Group) ToJSON() (string, error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", NewGroupToJSONError(err)
	}

	return string(jsonData), nil
}

func (a *Group) DecodeJSON(data string) error {
	err := json.Unmarshal([]byte(data), &a)
	if err != nil {
		return NewGroupDecodeJSONError(err)
	}

	return nil
}

func (a *Group) ParseMap(data map[string]interface{}) error {
	err := mapstructure.Decode(data, &a)
	if err != nil {
		return NewGroupParseMapError(err)
	}

	return nil
}

func (a *Group) GetPreloads() []string {
	return []string{
		"Users",
		"Owner",
	}
}
