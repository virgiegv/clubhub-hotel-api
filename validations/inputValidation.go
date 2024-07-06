package validations

import (
	"gopkg.in/go-playground/validator.v9"
)

func GetValidator() *validator.Validate {
	validate := validator.New()
	_ = validate.RegisterValidation("not-empty", notEmptyArray)
	return validate
}

func notEmptyArray(fl validator.FieldLevel) bool {
	return len(fl.Field().Interface().([]string)) != 0
}
