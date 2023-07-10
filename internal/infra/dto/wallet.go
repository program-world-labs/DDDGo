package dto

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
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

func (a *Wallet) Transform(i domain.IEntity) (IRepoEntity, error) {
	if err := copier.Copy(a, i); err != nil {
		return nil, NewWalletTransformError(err)
	}

	return a, nil
}

func (a *Wallet) BackToDomain() (domain.IEntity, error) {
	i := &entity.Wallet{}
	if err := copier.Copy(&i, a); err != nil {
		return nil, NewWalletBackToDomainError(err)
	}

	return i, nil
}

func (a *Wallet) BeforeUpdate(_ *gorm.DB) (err error) {
	a.UpdatedAt = time.Now()

	return
}
func (a *Wallet) BeforeCreate(_ *gorm.DB) (err error) {
	a.ID = uuid.New().String()
	a.UpdatedAt = time.Now()
	a.CreatedAt = time.Now()

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
		return "", NewWalletToJSONError(err)
	}

	return string(jsonData), nil
}

func (a *Wallet) DecodeJSON(data string) error {
	err := json.Unmarshal([]byte(data), &a)
	if err != nil {
		return NewWalletDecodeJSONError(err)
	}

	return nil
}