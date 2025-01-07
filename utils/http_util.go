package utils

import "github.com/go-playground/validator/v10"

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	}

	return "Unknown error"
}