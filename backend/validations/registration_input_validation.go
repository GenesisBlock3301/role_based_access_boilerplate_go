package validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

func RegistrationInputValidate(err error) []string {
	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()
		tagName := err.Tag()
		var errMsg []string
		switch tagName {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("%s is required", fieldName))
		case "email":
			errMsg = append(errMsg, fmt.Sprintf("%s must be a valid email address", fieldName))
		default:
			errMsg = append(errMsg, fmt.Sprintf("%s us invalid", fieldName))
		}
		return errMsg
	}
	return []string{}
}
