package entity

import (
	"time"

	"github.com/program-world-labs/DDDGo/internal/domain"
)

type Chain string

const (
	None     Chain = "None"
	Bitcoin  Chain = "Bitcoin"
	Ethereum Chain = "Ethereum"
	Polygon  Chain = "Polygon"
)

var _ domain.IEntity = (*Wallet)(nil)

type Wallet struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	Chain          Chain           `json:"chain"`
	Address        string          `json:"address"`
	UserID         string          `json:"userId"`
	WalletBalances []WalletBalance `json:"walletBalances"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      time.Time       `json:"deleted_at"`
}

func (a *Wallet) GetID() string {
	return a.ID
}

func (a *Wallet) SetID(id string) {
	a.ID = id
}
