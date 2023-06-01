package user

import "gitlab.com/demojira/template.git/internal/domain/user/entity"

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	EMail     string `json:"email"`
	Avatar    string `json:"avatar"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewUserResponse(model entity.User) UserResponse {
	return UserResponse{
		ID:       model.ID.String(),
		Username: model.Username,
		EMail:    model.EMail,
		Avatar:   model.Avatar,
	}
}
