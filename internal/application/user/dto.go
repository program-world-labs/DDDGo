package user

import "github.com/program-world-labs/DDDGo/internal/domain/entity"

type Output struct {
	ID          string `gorm:"primary_key;"`
	Username    string `json:"username"`
	EMail       string `json:"email"`
	DisplayName string `json:"display_name"`
	Avatar      string `json:"avatar"`
}

func NewOutput(e *entity.User) *Output {
	return &Output{
		ID:          e.ID,
		Username:    e.Username,
		EMail:       e.EMail,
		DisplayName: e.DisplayName,
		Avatar:      e.Avatar,
	}
}
