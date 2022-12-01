package helper

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func GetErrorData(err error) interface{} {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	var errors []FieldError
	for _, e := range errs {
		errors = append(errors, FieldError{
			Field: strings.ToLower(e.Field()),
			Error: e.ActualTag(),
		})
	}

	return errors
}
