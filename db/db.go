package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type Conn struct {
	conn *sqlx.DB
}

// DBコネクション作成
func NewDBConn() *sqlx.DB {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_OPTION"),
	)
	conn, err := sqlx.Connect(os.Getenv("DB_DRIVER"), connStr)

	if err != nil {
		log.Fatalf("error connecting database: %s", err.Error())
	}

	return conn
}

// DBコネクション切断
func (c *Conn) Close() {
	c.conn.Close()
}
