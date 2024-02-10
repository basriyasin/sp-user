// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"
)

type RepositoryInterface interface {
	// user profile mutation
	SaveProfile(ctx context.Context, user User) (id int64, err error)
	UpdateLoginCount(ctx context.Context, userID int64, loginCount int) (err error)
	UpdateUserByID(ctx context.Context, user User) error

	// user profile queries
	GetProfileByPhone(ctx context.Context, phone string) (user User, err error)
	GetProfileByID(ctx context.Context, id int64) (user User, err error)
	// end of user profile
}
