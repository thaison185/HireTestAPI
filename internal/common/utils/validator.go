package utils

import (
	"github.com/go-playground/validator/v10"

	"strings"
)

var Validate = validator.New()

func ParseValidationError(err error) string {
	if err == nil {
		return ""
	}

	validatorErr, ok := err.(validator.ValidationErrors)
	if !ok {
		return "validation error: " + err.Error() // Return the original error message if it's not a validation error
	}

	var errorMessages []string
	for _, fieldErr := range validatorErr {
		errorMessage := "validation failed on field '" + fieldErr.Field() + "' with tag '" + fieldErr.Tag() + "'"
		errorMessages = append(errorMessages, errorMessage)
	}

	return strings.Join(errorMessages, "; ")
}
