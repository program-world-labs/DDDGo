package dto

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/jinzhu/copier"
	"github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

var _ IRepoEntity = (*Role)(nil)

type Role struct {
	ID          string         `json:"id" gorm:"type:varchar(20);primary_key"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Permissions pq.StringArray `json:"permissions" gorm:"type:varchar(100)[]"`
	Users       []User         `json:"users" gorm:"many2many:user_roles;"`
	CreatedAt   time.Time      `json:"created_at" mapstructure:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" mapstructure:"updated_at"`
	DeletedAt   time.Time      `json:"deleted_at" mapstructure:"deleted_at" gorm:"index"`
}

func (a *Role) TableName() string {
	return "Roles"
}

func (a *Role) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRoleTransform, err)
	}

	return a, nil
}

func (a *Role) BackToDomain() (domain.IEntity, error) {
	i := &entity.Role{}
	if err := copier.Copy(i, a); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRoleBackToDomain, err)
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
		return "", domainerrors.Wrap(ErrorCodeRoleToJSON, err)
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
		return domainerrors.Wrap(ErrorCodeRoleDecodeJSON, err)
	}

	return nil
}

func (a *Role) ParseMap(data map[string]interface{}) (IRepoEntity, error) {
	err := ParseDateString(data)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeRoleParseMap, err)
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
		return nil, domainerrors.Wrap(ErrorCodeRoleParseMap, err)
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
