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

var _ IRepoEntity = (*WalletBalance)(nil)

type WalletBalance struct {
	ID         string          `json:"id" gorm:"primary_key"`
	WalletID   string          `json:"walletId" gorm:"index"`
	CurrencyID string          `json:"currencyId" gorm:"index"`
	Balance    uint            `json:"balance"`
	Decimal    uint            `json:"decimal"`
	CreatedAt  time.Time       `json:"created_at" mapstructure:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time       `json:"updated_at" mapstructure:"updated_at" gorm:"column:updated_at"`
	DeletedAt  *gorm.DeletedAt `json:"deleted_at" mapstructure:"deleted_at" gorm:"index;column:deleted_at"`
}

func (a *WalletBalance) TableName() string {
	return "WalletBalance"
}

func (a *WalletBalance) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeTransform, err)
	}

	return a, nil
}

func (a *WalletBalance) BackToDomain() (domain.IEntity, error) {
	i := &entity.Wallet{}
	if err := copier.Copy(i, a); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeBackToDomain, err)
	}

	if a.DeletedAt != nil {
		i.DeletedAt = a.DeletedAt.Time
	}

	return i, nil
}

func (a *WalletBalance) BeforeCreate() (err error) {
	a.ID, err = generateID()
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()
	a.DeletedAt = nil

	return
}

func (a *WalletBalance) BeforeUpdate() (err error) {
	a.UpdatedAt = time.Now()

	return
}

func (a *WalletBalance) GetID() string {
	return a.ID
}

func (a *WalletBalance) SetID(id string) {
	a.ID = id
}

func (a *WalletBalance) ToJSON() (string, error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", domainerrors.Wrap(ErrorCodeToJSON, err)
	}

	return string(jsonData), nil
}

func (a *WalletBalance) UnmarshalJSON(data []byte) error {
	type Alias WalletBalance

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

func (a *WalletBalance) ParseMap(data map[string]interface{}) (IRepoEntity, error) {
	err := ParseDateString(data)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeParseMap, err)
	}

	var info *WalletBalance
	err = mapstructure.Decode(data, &info)

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeParseMap, err)
	}

	return info, nil
}

func (a *WalletBalance) GetPreloads() []string {
	return []string{}
}

func (a *WalletBalance) GetListType() interface{} {
	entityType := reflect.TypeOf(Role{})
	sliceType := reflect.SliceOf(entityType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)

	return sliceValue.Interface()
}
