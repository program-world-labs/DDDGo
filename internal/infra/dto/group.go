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

var _ IRepoEntity = (*Group)(nil)

type Group struct {
	ID          string          `json:"id" gorm:"primary_key"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Users       []User          `json:"users" gorm:"foreignKey:GroupID"`
	Owner       *User           `json:"owner"`
	Metadata    string          `json:"metadata"`
	CreatedAt   time.Time       `json:"created_at" mapstructure:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time       `json:"updated_at" mapstructure:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" mapstructure:"deleted_at" gorm:"index;column:deleted_at"`
}

func (a *Group) TableName() string {
	return "Groups"
}

func (a *Group) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeGroupTransform, err)
	}

	return a, nil
}

func (a *Group) BackToDomain() (domain.IEntity, error) {
	i := &entity.Group{}
	if err := copier.Copy(i, a); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeGroupBackToDomain, err)
	}

	if a.DeletedAt != nil {
		i.DeletedAt = a.DeletedAt.Time
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
	a.DeletedAt = nil

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
		return "", domainerrors.Wrap(ErrorCodeGroupToJSON, err)
	}

	return string(jsonData), nil
}

func (a *Group) UnmarshalJSON(data []byte) error {
	type Alias Group

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return domainerrors.Wrap(ErrorCodeGroupDecodeJSON, err)
	}

	return nil
}

func (a *Group) ParseMap(data map[string]interface{}) (IRepoEntity, error) {
	err := ParseDateString(data)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeGroupParseMap, err)
	}

	var info *Group

	err = mapstructure.Decode(data, &info)

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeGroupParseMap, err)
	}

	return info, nil
}

func (a *Group) GetPreloads() []string {
	return []string{
		"Users",
		"Owner",
	}
}

func (a *Group) GetListType() interface{} {
	entityType := reflect.TypeOf(Role{})
	sliceType := reflect.SliceOf(entityType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)

	return sliceValue.Interface()
}
