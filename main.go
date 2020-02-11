package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/server"
)

func main() {
	// 環境変数(.env)読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	db.Init()
	defer db.Close()

	// APIサーバ起動
	server.Init()
}
