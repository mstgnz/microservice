package config

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate(structure interface{}) error {
	var validate *validator.Validate
	validate = validator.New()
	var errStr string
	var errSlc []error
	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(structure)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			errStr = err.Error()
		}
		for _, err := range err.(validator.ValidationErrors) {
			errStr = fmt.Sprintf("%s %s %s %s", err.Tag(), err.Param(), err.Field(), err.Type().String())
			errSlc = append(errSlc, errors.New(errStr))
			errStr = ""
		}
		// from here you can create your own error messages in whatever language you wish
		return errors.Join(errSlc...)
	}
	return nil
}
