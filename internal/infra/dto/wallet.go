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

type Chain string

const (
	None     Chain = "None"
	Bitcoin  Chain = "Bitcoin"
	Ethereum Chain = "Ethereum"
	Polygon  Chain = "Polygon"
)

var _ IRepoEntity = (*Wallet)(nil)

type Wallet struct {
	ID          string          `json:"id" gorm:"primary_key"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Chain       Chain           `json:"chain"`
	Address     string          `json:"address"`
	UserID      string          `json:"userId"`
	Amounts     []Amount        `json:"amounts" gorm:"foreignKey:WalletID"`
	CreatedAt   time.Time       `json:"created_at" mapstructure:"created_at" gorm:"column:created_at"`
	UpdatedAt   time.Time       `json:"updated_at" mapstructure:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at" mapstructure:"deleted_at" gorm:"index;column:deleted_at"`
}

func (a *Wallet) TableName() string {
	return "Wallet"
}

func (a *Wallet) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeWalletTransform, err)
	}

	return a, nil
}

func (a *Wallet) BackToDomain() (domain.IEntity, error) {
	i := &entity.Wallet{}
	if err := copier.Copy(i, a); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeWalletBackToDomain, err)
	}

	if a.DeletedAt != nil {
		i.DeletedAt = a.DeletedAt.Time
	}

	return i, nil
}

func (a *Wallet) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID, err = generateID()
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()
	a.DeletedAt = nil

	return
}

func (a *Wallet) GetID() string {
	return a.ID
}

func (a *Wallet) SetID(id string) {
	a.ID = id
}

func (a *Wallet) ToJSON() (string, error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", domainerrors.Wrap(ErrorCodeWalletToJSON, err)
	}

	return string(jsonData), nil
}

func (a *Wallet) UnmarshalJSON(data []byte) error {
	type Alias Wallet

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return domainerrors.Wrap(ErrorCodeWalletDecodeJSON, err)
	}

	return nil
}

func (a *Wallet) ParseMap(data map[string]interface{}) (IRepoEntity, error) {
	err := ParseDateString(data)
	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeWalletParseMap, err)
	}

	var info *Wallet
	err = mapstructure.Decode(data, &info)

	if err != nil {
		return nil, domainerrors.Wrap(ErrorCodeWalletParseMap, err)
	}

	return info, nil
}

func (a *Wallet) GetPreloads() []string {
	return []string{"Amounts"}
}

func (a *Wallet) GetListType() interface{} {
	entityType := reflect.TypeOf(Role{})
	sliceType := reflect.SliceOf(entityType)
	sliceValue := reflect.MakeSlice(sliceType, 0, 0)

	return sliceValue.Interface()
}
