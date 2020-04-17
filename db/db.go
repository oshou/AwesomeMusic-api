package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

// DBコネクション作成
func NewDBConn() *sqlx.DB {
	connStr := fmt.Sprintf(
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
	fmt.Println(connStr)
	conn, err := sqlx.Open(os.Getenv("DB_DRIVER"), connStr)

	if err != nil {
		log.Fatalf("error connecting database: %s", err.Error())
	}

	return conn
}
