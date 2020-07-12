// Package db is data access package
package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var pool *sqlx.DB

// Init is constructor for db
func Init() error {
	driver := os.Getenv("DB_DRIVER")
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	pool, err = sqlx.Open(driver, dsn)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func GetDB() *sqlx.DB {
	return pool
}

func Close() {
	if pool != nil {
		pool.Close()
	}
}
