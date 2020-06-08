package test

import (
	"database/sql"
	"path/filepath"

	testfixtures "gopkg.in/testfixtures.v2"
)

var (
	fdb      *sql.DB
	fixtures *testfixtures.Context
)

const (
	defaultPath = "../_fixtures"
)

func LoadFixtures() error {
	return LoadFixturesAt(defaultPath)
}

func LoadFixturesAt(path string) error {
	p, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	fdb, err = sql.Open("pgx", config.DB.DSN)

}
