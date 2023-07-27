package entity

import (
	"time"

	"github.com/program-world-labs/DDDGo/internal/domain"
)

var _ domain.IEntity = (*Currency)(nil)

type Currency struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Symbol         string          `json:"symbol"`
	WalletBalances []WalletBalance `json:"walletBalances"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      time.Time       `json:"deleted_at"`
}

func (a *Currency) GetID() string {
	return a.ID
}

func (a *Currency) SetID(id string) {
	a.ID = id
}
