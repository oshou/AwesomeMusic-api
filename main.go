package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/oshou/AwesomeMusic-api/repository"
	"github.com/oshou/AwesomeMusic-api/server"
)

func main() {
	// 環境変数(.env)読み込み
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	conn, err := sqlx.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?"+os.Getenv("DB_OPTION"))
	if err != nil {
		log.Fatalf("error connecting database: %s", err.Error())
	}
	userRepository := repository.NewMysqlUserRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userPresentor := presentor.userPresentor(userUsecase)

	// APIサーバ起動
	server.Init()
}
