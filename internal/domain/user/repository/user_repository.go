package repository

import (
	"gitlab.com/demojira/template.git/internal/domain"
	"gitlab.com/demojira/template.git/internal/domain/user/entity"
)

type UserRepository interface {
	domain.CRUDRepository[*entity.User]
}
