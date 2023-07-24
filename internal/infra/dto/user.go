package dto

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/entity"
)

var _ IRepoEntity = (*User)(nil)

type User struct {
	ID          string          `json:"id" gorm:"primary_key"`
	Username    string          `json:"username"`
	Password    string          `json:"password"`
	EMail       string          `json:"email"`
	DisplayName string          `json:"display_name"`
	Avatar      string          `json:"avatar"`
	Enabled     bool            `json:"enabled"`
	Roles       []Role          `json:"roles" gorm:"many2many:user_roles;"`
	Wallets     []Wallet        `json:"wallets" gorm:"foreignKey:UserID"`
	Group       Group           `json:"group"`
	GroupID     string          `json:"groupId" gorm:"index"`
	CreatedAt   time.Time       `json:"created_at" mapstructure:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time       `json:"updated_at" mapstructure:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" mapstructure:"deleted_at" gorm:"index;column:deleted_at"`
}

func (a *User) TableName() string {
	return "User"
}

func (a *User) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeUserTransform, err)
	}

	return a, nil
}

func (a *User) BackToDomain() (domain.IEntity, error) {
	var groupDelete time.Time
	if a.Group.DeletedAt != nil {
		groupDelete = a.Group.DeletedAt.Time
	}

	group := &entity.Group{
		ID:          a.GroupID,
		Name:        a.Group.Name,
		Description: a.Group.Description,
		Metadata:    a.Group.Metadata,
		CreatedAt:   a.Group.CreatedAt,
		UpdatedAt:   a.Group.UpdatedAt,
		DeletedAt:   groupDelete,
	}

	var deleteTime time.Time
	if a.DeletedAt != nil {
		deleteTime = a.DeletedAt.Time
	}

	i := &entity.User{
		ID:          a.ID,
		Username:    a.Username,
		Password:    a.Password,
		EMail:       a.EMail,
		DisplayName: a.DisplayName,
		Avatar:      a.Avatar,
		Group:       group,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
		DeletedAt:   deleteTime,
	}

	return i, nil
}

func (a *User) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID, err = generateID()
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()
	a.DeletedAt = nil

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
		return "", domainerrors.Wrap(ErrorCodeUserToJSON, err)
	}

	return string(jsonData), nil
}

func (a *User) UnmarshalJSON(data []byte) error {
	type Alias User

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return domainerrors.Wrap(ErrorCodeUserDecodeJSON, err)
	}

	return nil
}

func (a *User) ParseMap(data map[string]interface{}) (IRepoEntity, error) {
	err := ParseDateString(data)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeUserParseMap, err)
	}

	var info *User
	err = mapstructure.Decode(data, &info)

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeUserDecodeJSON, err)
	}

	return info, nil
}

func (a *User) GetPreloads() []string {
	return []string{"Roles", "Wallets", "Group"}
}

func (a *User) GetListType() interface{} {
	entityType := reflect.TypeOf(Role{})
	sliceType := reflect.SliceOf(entityType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)

	return sliceValue.Interface()
}
