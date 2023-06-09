package repository

import (
	"github.com/program-world-labs/DDDGo/internal/domain"
)

//go:generate mockgen -source=internal/domain/user/repository/user_repository.go -destination=tests/user/UserRepository_mock.go -package=user_test

type UserRepository interface {
	domain.ICRUDRepository
}
