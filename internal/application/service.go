package application

import (
	"github.com/program-world-labs/DDDGo/internal/application/group"
	"github.com/program-world-labs/DDDGo/internal/application/role"
	"github.com/program-world-labs/DDDGo/internal/application/user"
)

type Services struct {
	User  user.IService
	Role  role.IService
	Group group.IService
}
