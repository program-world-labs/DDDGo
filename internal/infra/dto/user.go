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

var _ IRepoEntity = (*User)(nil)

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

func (a *User) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, NewUserTransformError(err)
	}

	return a, nil
}

func (a *User) BackToDomain() (domain.IEntity, error) {
	i := &entity.User{}
	if err := copier.Copy(&i, a); err != nil {
		return nil, NewUserBackToDomainError(err)
	}

	return i, nil
}

func (a *User) BeforeUpdate(_ *gorm.DB) (err error) {
	a.UpdatedAt = time.Now()

	return
}
func (a *User) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID, err = generateID()
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()

	return
}

func (a *User) GetID() string {
	return a.ID
}

func (a *User) SetID(id string) {
	a.ID = id
}

func (a *User) ToJSON() (string, error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", NewUserToJSONError(err)
	}

	return string(jsonData), nil
}

func (a *User) DecodeJSON(data string) error {
	err := json.Unmarshal([]byte(data), &a)
	if err != nil {
		return NewUserDecodeJSONError(err)
	}

	return nil
}

func (a *User) ParseMap(data map[string]interface{}) error {
	if err := mapstructure.Decode(data, &a); err != nil {
		return NewUserParseMapError(err)
	}

	return nil
}

func (a *User) GetPreloads() []string {
	return []string{"Roles", "Wallets", "Group"}
}
