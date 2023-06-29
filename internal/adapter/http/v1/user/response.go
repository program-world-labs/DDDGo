package user

import (
	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
)

type Response struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	EMail     string `json:"email"`
	Avatar    string `json:"avatar"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewResponse(model application_user.Output) Response {
	return Response{
		ID:       model.ID,
		Username: model.Username,
		EMail:    model.EMail,
		Avatar:   model.Avatar,
	}
}
