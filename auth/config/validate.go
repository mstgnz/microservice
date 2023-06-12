package config

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Validate(structure interface{}) error {
	var validate *validator.Validate
	validate = validator.New()
	var errorStr string
	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(structure)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errorStr = err.Error()
		}
		for _, err := range err.(validator.ValidationErrors) {
			errorStr = errorStr + " " + err.Tag()
			errorStr = errorStr + " " + err.Field()
			errorStr = errorStr + " " + err.Type().String()
		}
		// from here you can create your own error messages in whatever language you wish
		return errors.New(errorStr)
	}
	return nil
}
