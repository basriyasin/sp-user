package handler

import (
	"fmt"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestValidate(t *testing.T) {
	type exampleRequired struct {
		Val *string `validate:"required"`
	}
	type exampleMin struct {
		Val string `validate:"min=3"`
	}
	type exampleMax struct {
		Val string `validate:"max=5"`
	}
	type exampleName struct {
		Val string `validate:"name"`
	}
	type examplePhone struct {
		Val string `validate:"phone"`
	}
	type examplePassword struct {
		Val string `validate:"password"`
	}
	type exampleUnexpected struct {
		Val string `validate:"unknown"`
	}

	validate.RegisterAlias("unknown", "password")

	test := []struct {
		name string
		args interface{}
		err  error
	}{
		{
			name: "err required",
			args: exampleRequired{},
			err:  fmt.Errorf("'Val' is required"),
		},
		{
			name: "err min=3",
			args: exampleMin{Val: "1"},
			err:  fmt.Errorf("'Val' should be minimum 3"),
		},
		{
			name: "err max=5",
			args: exampleMax{Val: "123456"},
			err:  fmt.Errorf("'Val' should be maximum 5"),
		},
		{
			name: "err name",
			args: exampleName{Val: "123"},
			err:  fmt.Errorf("'Val' should have at lease have %d and max %d alpha character", nameMinLength, nameMaxLength),
		},
		{
			name: "err phone",
			args: examplePhone{Val: "+6082211223344"},
			err:  fmt.Errorf("'Val' should have start with %s, have min %d and max %d character ", phonePrefix, phoneMinLength, phoneMaxLength),
		},
		{
			name: "err password",
			args: examplePassword{Val: "abcd"},
			err:  fmt.Errorf("'Val' should contain at lease 1 lower case, 1 upper case and 1 special character"),
		},
		{
			name: "unexpeced err",
			args: exampleUnexpected{Val: "unex"},
			err:  fmt.Errorf("'Val' error"),
		},
		{
			name: "err password",
			args: examplePassword{Val: "Aa12!@"},
			err:  nil,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := Validate(tt.args)
			assert.Equal(t, tt.err, got)
		})
	}
}
