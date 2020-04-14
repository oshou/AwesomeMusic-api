package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/interactor"
	"github.com/oshou/AwesomeMusic-api/presenter/middleware"
	"github.com/oshou/AwesomeMusic-api/presenter/router"
)

func main() {
	// 環境変数(.env)読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	conn := db.NewDBConn()
	i := interactor.NewInteractor(conn)
	e := gin.Default()
	h := i.NewAppHandler()
	router.NewRouter(e, h)
	middleware.NewMiddleware(e)
	if err := e.Run(":" + os.Getenv("API_SERVER_PORT")); err != nil {
		log.Fatalln(err)
	}
}
