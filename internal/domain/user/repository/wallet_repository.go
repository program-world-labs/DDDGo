package repository

import (
	"github.com/program-world-labs/DDDGo/internal/domain"
)

type WalletRepository interface {
	domain.ICRUDRepository
}
