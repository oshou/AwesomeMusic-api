package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type DB struct {
	conn *sqlx.DB
}

func NewDB() *DB {
	db, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?"+os.Getenv("DB_OPTION"))
	if err != nil {
		log.Fatalf("error connecting database: %s", err.Error())
	}
	return &DB{
		conn: db,
	}
}

func (db *DB) DBConn() *sqlx.DB {
	return db.conn
}

func (db *DB) Close() {
	db.conn.Close()
}

// type DataStore interface {
// 	Get() DataStore
// 	Close()
// }
//
// type DB struct {
// 	conn *sqlx.DB
// }
//
// func NewDB() DataStore {
// 	db, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?"+os.Getenv("DB_OPTION"))
// 	if err != nil {
// 		log.Fatalf("error connecting database: %s", err.Error())
// 	}
//
// 	return &DB{
// 		conn: db,
// 	}
// }
//
// // 作成済DBコネクションの取得
// func (db *DB) Get() DataStore {
// 	return db
// }
//
// // 作成済DBコネクションの取得
// func (db *DB) Close() {
// 	db.conn.Close()
// }
