package user

import "github.com/program-world-labs/DDDGo/internal/domain/user/entity"

type Response struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	EMail     string `json:"email"`
	Avatar    string `json:"avatar"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewResponse(model entity.User) Response {
	return Response{
		ID:       model.ID.String(),
		Username: model.Username,
		EMail:    model.EMail,
		Avatar:   model.Avatar,
	}
}
