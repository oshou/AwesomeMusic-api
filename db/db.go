package repository

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var conn Repositry

type MysqlUserRepositry interface {
	Conn *sqlx.DB
}

// DBコネクション作成
func NewMysqlUserRepository() repository. {
	conn, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?"+os.Getenv("DB_OPTION"))
	if err != nil {
		log.Fatalf("error connecting database: %s", err.Error())
	}
	return DB{
		conn: conn,
	}
}

// 作成済DBコネクションの取得
func (db *DB) DBConn() Repositry {
	return db.conn
}

// DBコネクション切断
func Close() {
	conn.Close()
}
