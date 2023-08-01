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

var _ IRepoEntity = (*Role)(nil)

type Role struct {
	ID          string          `json:"id" gorm:"type:varchar(20);primary_key" firestore:"id"`
	Name        string          `json:"name" firestore:"name"`
	Description string          `json:"description" firestore:"description"`
	Permissions []string        `json:"permissions" gorm:"type:json" firestore:"permissions"`
	Users       []User          `json:"users" gorm:"many2many:user_roles;" firestore:"users"`
	CreatedAt   time.Time       `json:"created_at" mapstructure:"created_at" gorm:"column:created_at" firestore:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" mapstructure:"updated_at" gorm:"column:updated_at" firestore:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" mapstructure:"deleted_at" gorm:"index;column:deleted_at" firestore:"deleted_at"`
}

func (a *Role) TableName() string {
	return "Roles"
}

func (a *Role) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeTransform, err)
	}

	return a, nil
}

func (a *Role) BackToDomain() (domain.IEntity, error) {
	i := &entity.Role{}
	if err := copier.Copy(i, a); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeBackToDomain, err)
	}

	if a.DeletedAt != nil {
		i.DeletedAt = a.DeletedAt.Time
	}

	return i, nil
}

func (a *Role) BeforeUpdate() (err error) {
	a.UpdatedAt = time.Now()

	return
}

func (a *Role) BeforeCreate() (err error) {
	a.ID, err = generateID()
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()
	a.DeletedAt = nil

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
		return "", domainerrors.Wrap(ErrorCodeToJSON, err)
	}

	return string(jsonData), nil
}

func (a *Role) UnmarshalJSON(data []byte) error {
	type Alias Role

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return domainerrors.Wrap(ErrorCodeDecodeJSON, err)
	}

	return nil
}

func (a *Role) ParseMap(data map[string]interface{}) (IRepoEntity, error) {
	err := ParseDateString(data)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeParseMap, err)
	}

	var info *Role
	// Permissions is a slice of string, so we need to decode it manually, data like {read:all,write:all}
	// permission, ok := data["permissions"].(string)
	// if !ok {
	// 	return nil, NewRoleParseMapError(nil)
	// }

	// s := strings.Trim(permission, "{}") // 删除开头和结尾的大括号
	// result := strings.Split(s, ",")     // 以逗号为分割符，分割字符串
	// data["permissions"] = result

	err = mapstructure.Decode(data, &info)

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeParseMap, err)
	}

	return info, nil
}

func (a *Role) GetPreloads() []string {
	return []string{
		"Users",
	}
}

func (a *Role) GetListType() interface{} {
	entityType := reflect.TypeOf(Role{})
	sliceType := reflect.SliceOf(entityType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)

	return sliceValue.Interface()
}
