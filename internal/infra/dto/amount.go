package dto

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/domainerrors"
	"github.com/program-world-labs/DDDGo/internal/domain/entity"
)

var _ IRepoEntity = (*Amount)(nil)

type Amount struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Currency  string    `json:"currency"`
	Icon      string    `json:"icon"`
	Balance   uint      `json:"balance"`
	Decimal   uint      `json:"decimal"`
	WalletID  string    `json:"walletId" gorm:"index"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" gorm:"index"`
}

func (a *Amount) TableName() string {
	return "Amounts"
}

func (a *Amount) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeAmountTransform, err)
	}

	return a, nil
}

func (a *Amount) BackToDomain() (domain.IEntity, error) {
	i := &entity.Amount{}
	if err := copier.Copy(&i, a); err != nil {
		return nil, domainerrors.Wrap(ErrorCodeAmountBackToDomain, err)
	}

	return i, nil
}

func (a *Amount) BeforeUpdate(_ *gorm.DB) (err error) {
	a.UpdatedAt = time.Now()

	return
}

func (a *Amount) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID, err = generateID()
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()

	return
}

func (a *Amount) GetID() string {
	return a.ID
}

func (a *Amount) SetID(id string) {
	a.ID = id
}

func (a *Amount) ToJSON() (string, error) {
	jsonData, err := json.Marshal(a)
	if err != nil {
		return "", domainerrors.Wrap(ErrorCodeAmountToJSON, err)
	}

	return string(jsonData), nil
}

func (a *Amount) DecodeJSON(data string) error {
	err := json.Unmarshal([]byte(data), &a)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeAmountDecodeJSON, err)
	}

	return nil
}

func (a *Amount) ParseMap(data map[string]interface{}) error {
	err := mapstructure.Decode(data, &a)
	if err != nil {
		return domainerrors.Wrap(ErrorCodeAmountParseMap, err)
	}

	return nil
}

func (a *Amount) GetPreloads() []string {
	return []string{}
}
