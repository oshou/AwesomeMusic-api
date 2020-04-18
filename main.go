package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/injector"
)

func main() {
	// Load Environment
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	// set DBConnection
	db, err := db.NewDB()
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// Routing
	i := injector.NewInjector(db)
	r := i.NewRouter()

	if err := r.Run(":" + os.Getenv("API_SERVER_PORT")); err != nil {
		log.Fatalln(err)
	}
}
