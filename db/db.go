// Package db is data access package
package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// NewDB is constructor for db
func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open(
		os.Getenv("DB_DRIVER"),
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}
