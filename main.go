package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/injector"
	"github.com/oshou/AwesomeMusic-api/presenter/middleware"
	"github.com/oshou/AwesomeMusic-api/presenter/router"
)

func main() {
	// 環境変数(.env)読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	// DBConnection作成
	conn := db.NewDBConn()
	defer conn.Close()

	// Routing
	i := injector.NewInjector(conn)
	handler := i.NewAppHandler()
	engine := gin.Default()
	middleware.NewMiddleware(engine)
	router.NewRouter(engine, handler)

	if err := engine.Run(":" + os.Getenv("API_SERVER_PORT")); err != nil {
		log.Fatalln(err)
	}
}
