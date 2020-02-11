package db

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	conn *sqlx.DB
	err  error
)

// DBコネクション作成
func Init() {
	conn, err = sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?"+os.Getenv("DB_OPTION"))
	if err != nil {
		log.Fatalf("error connecting database: %s", err.Error())
	}
}

// 作成済DBコネクションの取得
func DBConn() *sqlx.DB {
	return conn
}

// DBコネクション切断
func Close() {
	conn.Close()
}
