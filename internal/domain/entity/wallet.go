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
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Chain       Chain     `json:"chain"`
	Address     string    `json:"address"`
	UserID      string    `json:"userId"`
	Amounts     []Amount  `json:"amounts"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt"`
}

func (a *Wallet) GetID() string {
	return a.ID
}

func (a *Wallet) SetID(id string) {
	a.ID = id
}
