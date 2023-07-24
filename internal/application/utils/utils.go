package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	ErrValidation = errors.New("validation failed")
)

func HandleValidationError(structName string, validateErrors validator.ValidationErrors) error {
	tagErrors := map[string]func(string) string{
		"required": func(field string) string {
			return structName + " " + field + " required"
		},
		"lte": func(field string) string {
			return structName + " " + field + " exceeds max length"
		},
		"alphanum": func(field string) string {
			return structName + " " + field + " invalid format"
		},
		"custom_permission": func(field string) string {
			return structName + " " + field + " invalid permission format"
		},
		"oneof": func(field string) string {
			return structName + " " + field + " invalid sort field"
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
