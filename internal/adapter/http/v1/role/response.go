package role

import (
	"time"

	"github.com/program-world-labs/DDDGo/internal/adapter/http/v1/user"
	application_role "github.com/program-world-labs/DDDGo/internal/application/role"
)

type Response struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Permissions []string        `json:"permissions"`
	Users       []user.Response `json:"users"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

func NewResponse(model *application_role.Output) Response {
	userList := make([]user.Response, len(model.Users))
	for i, v := range model.Users {
		userList[i] = user.NewResponse(v)
	}

	return Response{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		Permissions: model.Permissions,
		Users:       userList,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}
