package dto

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

var _ entity.IEntity = (*Amount)(nil)

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

func (a *Amount) Transform(i domain.IEntity) (entity.IEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, NewAmountTransformError(err)
	}

	return a, nil
}

func (a *Amount) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID = uuid.New().String()

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
		return "", NewAmountToJSONError(err)
	}

	return string(jsonData), nil
}

func (a *Amount) DecodeJSON(data string) error {
	err := json.Unmarshal([]byte(data), &a)
	if err != nil {
		return NewAmountDecodeJSONError(err)
	}

	return nil
}
