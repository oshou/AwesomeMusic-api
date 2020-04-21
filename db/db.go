// Package db is data access package
package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// NewDB is constructor for db
func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}
