package role

import (
	"time"

	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

type CreatedInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

func (c *CreatedInput) ToEntity() *entity.Role {
	return &entity.Role{
		Name:        c.Name,
		Description: c.Description,
		Permissions: c.Permissions,
	}
}

type AssignedInput struct {
	ID     string   `json:"id"`
	UserID []string `json:"userId"`
	Assign bool     `json:"assign"`
}

type UpdatedInput struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

type ListGotInput struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Output struct {
	ID          string                    `json:"id"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Permissions []string                  `json:"permissions"`
	Users       []application_user.Output `json:"users"`
	CreatedAt   time.Time                 `json:"createdAt"`
	UpdatedAt   time.Time                 `json:"updatedAt"`
	DeletedAt   time.Time                 `json:"deletedAt"`
}

func NewOutput(e *entity.Role) *Output {
	return &Output{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Permissions: e.Permissions,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   e.DeletedAt,
	}
}
