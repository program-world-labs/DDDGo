package entity

import "time"

type WalletBalance struct {
	ID         string    `json:"id"`
	WalletID   string    `json:"walletId"`
	CurrencyID string    `json:"currencyId"`
	Balance    uint      `json:"balance"`
	Decimal    uint      `json:"decimal"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
