package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

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

func (a *Amount) Self() interface{} {
	return a
}
