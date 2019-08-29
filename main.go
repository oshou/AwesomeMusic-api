package main

import (
	"log"

	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/server"
	"github.com/joho/godotenv"
)

func main() {
	// 環境変数(.env)読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
	db.Init()

	// 終了時に遅延実行でDBコネクション切断されるよう設定
	defer db.Close()

	// APIサーバ起動
	server.Init()
}
