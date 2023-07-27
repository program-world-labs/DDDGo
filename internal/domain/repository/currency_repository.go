package repository

import (
	"github.com/program-world-labs/DDDGo/internal/domain"
)

type CurrencyRepository interface {
	domain.ICRUDRepository
}
