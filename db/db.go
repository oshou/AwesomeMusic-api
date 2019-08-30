package db

import (
	"log"
	"os"

	"github.com/oshou/AwesomeMusic-api/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

// DBコネクション作成
func Init() {
	db, err = gorm.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+")/"+os.Getenv("DB_NAME")+"?"+os.Getenv("DB_OPTION"))
	// db.LogMode(true)
	if err != nil {
		log.Fatalf("Unable to connect DB : %s", err.Error())
	}
}

// 作成済DBコネクションの取得
func GetDBConn() *gorm.DB {
	return db
}

// DBコネクション切断
func Close() {
	db.Close()
}
