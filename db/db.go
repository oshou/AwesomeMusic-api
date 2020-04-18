// Package db is data access package
package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

type db struct {
	conn *sqlx.DB
}

// NewDB is constructor for db
func NewDB() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		// - for MySQL
		//"%s:%s@tcp(%s:%s)/%s?%s",
		// - for Postgres
		"postgres://%s:%s@%s:%s/%s?%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_OPTION"),
	)
	db, err := sqlx.Open(os.Getenv("DB_DRIVER"), dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *db) Close() error {
	err := db.conn.Close()
	return err
}
