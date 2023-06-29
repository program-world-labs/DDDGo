package role

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type CreatedRequest struct {
	Name        string   `json:"name" validate:"required,lte=30" example:"admin"`
	Description string   `json:"description" validate:"required,lte=200" example:"this is for admin role"`
	Permissions []string `json:"permissions" validate:"required,dive,custom_permission" example:"read:all,write:all"`
}

func CustomPermission(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^(read|write|delete):.*$`)

	return re.MatchString(fl.Field().String())
}
