package repository

import (
	"github.com/program-world-labs/DDDGo/internal/domain"
)

type UserRepository interface {
	domain.ICRUDRepository
}
