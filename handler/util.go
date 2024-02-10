package handler

import (
	"regexp"
	"strings"
	"unicode"
)

const (
	phonePrefix = "+62"

	// name length
	nameMinLength = 3
	nameMaxLength = 60

	// password length
	passwordMinLength = 6
	passwordMaxLenght = 64

	// phone length
	phoneMinLength = 10
	phoneMaxLength = 13

	// default date format
	DateTimeFormat string = "2015-09-02 08:04:00"
)
const ()

var (
	// regex for validating user name
	nameRegex = regexp.MustCompile(`([a-zA-Z ]){3,64}`)
)

// Validate the user's name and return true if the name passes validation
func isValidName(name string) bool {
	length := len(name)
	if length < nameMinLength || length > nameMaxLength {
		return false
	}

	return nameRegex.MatchString(name)
}

// Validate the user's phone number and return true if the phone number passes validation.
func isValidPhone(phone string) bool {
	length := len(phone)
	if length < phoneMinLength+1 || length > phoneMaxLength+1 {
		return false
	}
	for _, char := range phone {
		switch {
		case unicode.IsNumber(char), char == '+':
			continue
		default:
			return false
		}
	}

	return strings.HasPrefix(phone, phonePrefix)
}

// Validate the user's password and return true if the password passes validation.
func isValidPassword(pass string) bool {
	length := len(pass)
	if length < passwordMinLength || length > passwordMaxLenght {
		return false
	}

	var (
		hasUpper, hasLower, hasNumber, hasSpecial bool
	)
	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}
