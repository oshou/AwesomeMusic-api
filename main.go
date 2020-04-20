package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/injector"
)

func main() {
	// Load Environment
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading .env file: %+v\n", err)
		log.Fatalln()
	}

	// set DBConnection
	db, err := db.NewDB()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	}()

	//err = db.Ping()
	//fmt.Println(err)

	// Routing
	i := injector.NewInjector(db)
	r := i.NewRouter()

	if err := r.Run(":" + os.Getenv("API_SERVER_PORT")); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
