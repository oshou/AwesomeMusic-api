// Package db is data access package
package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/oshou/AwesomeMusic-api/config"
)

type IDB interface {
	Close()
}

type db struct {
	Pool *sqlx.DB
}

var _ IDB = &db{}

func NewDB(c config.IConfig) (*db, error) {
	dsn := c.GetDSN()
	driver := c.GetDriver()
	pool, err := sqlx.Open(driver, dsn)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &db{pool}, nil
}

func (db *db) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
