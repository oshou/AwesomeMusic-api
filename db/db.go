// Package db is data access package
package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type db struct {
	conn *sqlx.DB
}

// NewDB is constructor for db
func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}

func (db *db) Close() error {
	err := db.conn.Close()
	return errors.WithStack(err)
}
