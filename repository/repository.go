// This file contains the repository implementation layer.
package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

type NewRepositoryOptions struct {
	Db *sql.DB
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	return &Repository{
		Db: opts.Db,
	}
}
