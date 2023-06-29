package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/infra/base/entity"
)

type Chain string

const (
	None     Chain = "None"
	Bitcoin  Chain = "Bitcoin"
	Ethereum Chain = "Ethereum"
	Polygon  Chain = "Polygon"
)

var _ entity.IEntity = (*Wallet)(nil)

type Wallet struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Chain       Chain     `json:"chain"`
	Address     string    `json:"address"`
	UserID      string    `json:"userId"`
	Amounts     []Amount  `json:"amounts" gorm:"foreignKey:WalletID"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt" gorm:"index"`
}

func (a *Wallet) TableName() string {
	return "Wallets"
}

func (a *Wallet) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID = uuid.New().String()

	return
}

func (a *Wallet) GetID() string {
	return a.ID
}

func (a *Wallet) SetID(id string) {
	a.ID = id
}

func (a *Wallet) Self() interface{} {
	return a
}
