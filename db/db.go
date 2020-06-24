// Package db is data access package
package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	Pool    *sqlx.DB
	maxconn = 10
)

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
	fmt.Println("dsn:", dsn)

	Pool, err := sqlx.Open(driver, dsn)
	if err != nil {
		return errors.WithStack(err)
	}
	Pool.SetMaxIdleConns(maxconn)

	return nil
}

func Close() {
	if Pool != nil {
		Pool.Close()
	}
}
