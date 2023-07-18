package repository

import (
	"github.com/program-world-labs/DDDGo/internal/domain"
)

type GroupRepository interface {
	domain.ICRUDRepository
}
