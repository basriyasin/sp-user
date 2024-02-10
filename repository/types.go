// This file contains types that are used in the repository layer.
package repository

import (
	"database/sql"
	"time"
)

type (

	// Considering the simplicity of the project architecture, authentication, profile, and user phone number will be put into one table to maintain simplicity.
	// For larger projects, it is recommended to separate authentication, user profile, and user phone into separate tables to increase flexibility in handling future scenarios, such as login with multiple authorization methods, a single user with multiple phone numbers, etc.
	User struct {
		ID         int64        `json:"id"`
		Phone      string       `json:"phone"        validate:"required,phone"`
		Name       string       `json:"name"         validate:"required,min=3,max=60"`
		Password   string       `json:"password"     validate:"required,min=6,max=64,password"`
		LoginCount int          `json:"login_count"`
		CreatedAt  time.Time    `json:"created_at"`
		UpdatedAt  sql.NullTime `json:"updated_at"`
	}
)
