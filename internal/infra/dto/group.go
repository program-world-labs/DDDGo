package dto

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
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

	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()

	return a, nil
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
