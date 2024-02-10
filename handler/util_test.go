package handler

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestIsValidName(t *testing.T) {
	test := []struct {
		name    string
		isValid bool
		args    string
	}{
		{
			name:    "empty name",
			args:    "",
			isValid: false,
		},
		{
			name:    "number only",
			args:    "12345",
			isValid: false,
		},
		{
			name:    "less than 3 character",
			args:    "ab",
			isValid: false,
		},
		{
			name:    "61 character",
			args:    "abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefa",
			isValid: false,
		},
		{
			name:    "have less than 3 alpha character",
			args:    "ab345",
			isValid: false,
		},
		{
			name:    "min length",
			args:    "abc",
			isValid: true,
		},
		{
			name:    "max length",
			args:    "abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdef",
			isValid: true,
		},
		{
			name:    "contain special character",
			args:    "Mr. Junior",
			isValid: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidName(tt.args)
			assert.Equal(t, tt.isValid, got)
		})
	}
}

func TestIsValidPhone(t *testing.T) {
	test := []struct {
		name    string
		isValid bool
		args    string
	}{
		{
			name:    "empty phone",
			args:    "",
			isValid: false,
		},
		{
			name:    "9 character",
			args:    "+62811223",
			isValid: false,
		},
		{
			name:    "14 character",
			args:    "+62811223344556",
			isValid: false,
		},
		{
			name:    "contain non number",
			args:    "+6281x2233445",
			isValid: false,
		},
		{
			name:    "not start with +62",
			args:    "+61811223344",
			isValid: false,
		},
		{
			name:    "correct phone",
			args:    "+62811223344",
			isValid: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidPhone(tt.args)
			assert.Equal(t, tt.isValid, got)
		})
	}
}

func TestIsValidPPassword(t *testing.T) {
	test := []struct {
		name    string
		isValid bool
		args    string
	}{
		{
			name:    "empty password",
			args:    "",
			isValid: false,
		},
		{
			name:    "5 character",
			args:    "Aa1@s",
			isValid: false,
		},
		{
			name:    "65 character",
			args:    "Aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1@",
			isValid: false,
		},
		{
			name:    "not have lower case",
			args:    "AAA12!@",
			isValid: false,
		},
		{
			name:    "not have upper case",
			args:    "aaa12!@",
			isValid: false,
		},
		{
			name:    "not have number",
			args:    "aaaAA!@",
			isValid: false,
		},
		{
			name:    "not have special character",
			args:    "aaaAA12",
			isValid: false,
		},
		{
			name:    "min password",
			args:    "Ab12!@",
			isValid: true,
		},
		{
			name:    "max password",
			args:    "Aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa1@",
			isValid: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidPassword(tt.args)
			assert.Equal(t, tt.isValid, got)
		})
	}
}
