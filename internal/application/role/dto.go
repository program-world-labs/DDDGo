package role

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"

	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
	"github.com/program-world-labs/DDDGo/internal/domain/user/entity"
)

type CreatedInput struct {
	Name        string `json:"name" validate:"required,lte=30,alphanum"`
	Description string `json:"description" validate:"required,lte=200"`
	Permissions string `json:"permissions" validate:"required,custom_permission"`
}

func CustomPermission(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^((read|write|delete):[a-z]+,)*((read|write|delete):[a-z]+)$`)

	return re.MatchString(fl.Field().String())
}

func handleValidationError(validateErrors validator.ValidationErrors) error {
	tagErrors := map[string]func(string) string{
		"required": func(field string) string {
			return "role " + field + " required"
		},
		"lte": func(field string) string {
			return "role " + field + " exceeds max length"
		},
		"alphanum": func(field string) string {
			return "role " + field + " invalid format"
		},
		"custom_permission": func(field string) string {
			return "role " + field + " invalid permission format"
		},
	}

	var errorMessages []string

	for _, err := range validateErrors {
		if specificError, ok := tagErrors[err.Tag()]; ok {
			errorMessages = append(errorMessages, specificError(err.Field()))
		} else {
			errorMessages = append(errorMessages, err.Error())
		}
	}

	return fmt.Errorf("%w: %s", ErrValidation, strings.Join(errorMessages, ", "))
}

func (c *CreatedInput) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("custom_permission", CustomPermission)

	if err != nil {
		return err
	}

	// 提取錯誤訊息
	err = validate.Struct(c)
	if err != nil {
		var e validator.ValidationErrors
		if errors.As(err, &e) {
			return handleValidationError(e)
		}

		return err
	}

	return nil
}

func (c *CreatedInput) ToEntity() *entity.Role {
	return &entity.Role{
		Name:        c.Name,
		Description: c.Description,
		Permissions: strings.Split(c.Permissions, ","),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
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
