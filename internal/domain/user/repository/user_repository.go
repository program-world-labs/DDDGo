package repository

import (
	"github.com/program-world-labs/DDDGo/internal/domain"
	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

type UserRepository interface {
	domain.ICRUDRepository[*entity.User]
}
