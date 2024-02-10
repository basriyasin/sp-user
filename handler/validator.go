package handler

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("name", validateName)
	validate.RegisterValidation("phone", validatePhone)
	validate.RegisterValidation("password", validatePassword)
}

// Validate the struct with the given tags on its fields and return an error for each field.
func Validate(s interface{}) (err error) {
	err = validate.Struct(s)
	if err == nil {
		return nil
	}

	//replace error message with custom error message
	verr, _ := err.(validator.ValidationErrors)
	var msg []string
	for _, e := range verr {
		switch e.Tag() {
		case "required":
			msg = append(msg, fmt.Sprintf("'%s' is required", e.Field()))
		case "min":
			msg = append(msg, fmt.Sprintf("'%s' should be minimum %s", e.Field(), e.Param()))
		case "max":
			msg = append(msg, fmt.Sprintf("'%s' should be maximum %s", e.Field(), e.Param()))
		case "name":
			msg = append(msg, fmt.Sprintf("'%s' should have at lease have %d and max %d alpha character", e.Field(), nameMinLength, nameMaxLength))
		case "phone":
			msg = append(msg, fmt.Sprintf("'%s' should have start with %s, have min %d and max %d character ", e.Field(), phonePrefix, phoneMinLength, phoneMaxLength))
		case "password":
			msg = append(msg, fmt.Sprintf("'%s' should contain at lease 1 lower case, 1 upper case and 1 special character", e.Field()))
		default:
			msg = append(msg, fmt.Sprintf("'%s' error", e.Field()))
		}
	}

	return errors.New(strings.Join(msg, " | "))
}

// custom validator for 'name' tag
func validateName(fl validator.FieldLevel) bool {
	return isValidName(fl.Field().String())
}

// custom validator for 'phone' tag
func validatePhone(fl validator.FieldLevel) bool {
	return isValidPhone(fl.Field().String())
}

// custom validator for 'passowrd' tag
func validatePassword(fl validator.FieldLevel) bool {
	return isValidPassword(fl.Field().String())
}
