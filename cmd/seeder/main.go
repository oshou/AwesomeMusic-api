package main

import (
	"database/sql"
	"fmt"

	testfixtures "github.com/go-testfixtures/testfixtures/v3"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/config"
	"github.com/oshou/AwesomeMusic-api/log"
)

var (
	configFilePath = "./config/config.yml"
	seedFilePath   = "./_db/seed"
	db             *sql.DB
	fixtures       *testfixtures.Loader
)

func main() {

	// Logger
	log.Init()
	defer log.Logger.Sync()
	log.Logger.Info("set logger")

	// Config
	conf, err := config.NewConfig(configFilePath)
	if err != nil {
		log.Logger.Fatal("failed to initialize config", zap.Error(err))
	}
	log.Logger.Info("set config")

	// DBConnection
	fmt.Println("dsn:", conf.GetDSN())
	db, err := sql.Open(conf.GetDriver(), conf.GetDSN())
	if err != nil {
		log.Logger.Fatal("failed to initialize db", zap.Error(err))
	}
	defer db.Close()
	log.Logger.Info("set db connection")

	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(seedFilePath),
	)
	if err != nil {
		log.Logger.Fatal("failed to create seeder", zap.Error(err))
	}
	err = fixtures.EnsureTestDatabase()
	if err != nil {
		log.Logger.Fatal("failed to seed dba", zap.Error(err))
	}

	if err := fixtures.Load(); err != nil {
		log.Logger.Fatal("failed to seed db", zap.Error(err))
	}
	log.Logger.Info("success to seed db")
}
