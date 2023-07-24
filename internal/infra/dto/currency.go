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

var _ IRepoEntity = (*Currency)(nil)

type Currency struct {
	ID             string `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	WalletBalances []WalletBalance
	CreatedAt      time.Time       `json:"created_at" mapstructure:"created_at" gorm:"column:created_at"`
	UpdatedAt      time.Time       `json:"updated_at" mapstructure:"updated_at" gorm:"column:updated_at"`
	DeletedAt      *gorm.DeletedAt `json:"deleted_at" mapstructure:"deleted_at" gorm:"index;column:deleted_at"`
}

func (a *Currency) TableName() string {
	return "Currency"
}

func (a *Currency) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCurrencyTransform, err)
	}

	return a, nil
}

func (a *Currency) BackToDomain() (domain.IEntity, error) {
	i := &entity.Currency{}
	if err := copier.Copy(i, a); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCurrencyBackToDomain, err)
	}

	if a.DeletedAt != nil {
		i.DeletedAt = a.DeletedAt.Time
	}

	return i, nil
}

func (a *Currency) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID, err = generateID()
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()
	a.DeletedAt = nil

	return
}

func (a *Currency) GetID() string {
	return a.ID
}

func (a *Currency) SetID(id string) {
	a.ID = id
}

func (a *Currency) ToJSON() (string, error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", domainerrors.Wrap(ErrorCodeCurrencyToJSON, err)
	}

	return string(jsonData), nil
}

func (a *Currency) UnmarshalJSON(data []byte) error {
	type Alias Currency

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return domainerrors.Wrap(ErrorCodeCurrencyDecodeJSON, err)
	}

	return nil
}

func (a *Currency) ParseMap(data map[string]interface{}) (IRepoEntity, error) {
	err := ParseDateString(data)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCurrencyParseMap, err)
	}

	var info *Currency

	err = mapstructure.Decode(data, &info)

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeCurrencyParseMap, err)
	}

	return info, nil
}

func (a *Currency) GetPreloads() []string {
	return []string{}
}

func (a *Currency) GetListType() interface{} {
	entityType := reflect.TypeOf(Role{})
	sliceType := reflect.SliceOf(entityType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)

	return sliceValue.Interface()
}
