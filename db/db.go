package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type Conn struct {
	conn *sqlx.DB
}

// DBコネクション作成
func NewDBConn() *sqlx.DB {
	conn, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?"+os.Getenv("DB_OPTION"))
	if err != nil {
		log.Fatalf("error connecting database: %s", err.Error())
	}
	return conn
}

// 作成済DBコネクションの取得
func (c *Conn) DBConn() *sqlx.DB {
	return c.conn
}

// DBコネクション切断
func (c *Conn) Close() {
	c.conn.Close()
}
