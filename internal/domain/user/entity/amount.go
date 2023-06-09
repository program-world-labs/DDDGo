package entity

import (
	"time"

	"github.com/program-world-labs/DDDGo/internal/domain"
)

var _ domain.IEntity = (*Amount)(nil)

type Amount struct {
	ID        string    `json:"id"`
	Currency  string    `json:"currency"`
	Icon      string    `json:"icon"`
	Balance   uint      `json:"balance"`
	Decimal   uint      `json:"decimal"`
	WalletID  string    `json:"walletId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func (a *Amount) GetID() string {
	return a.ID
}

func (a *Amount) SetID(id string) {
	a.ID = id
}
