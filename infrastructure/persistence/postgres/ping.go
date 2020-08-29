package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
	"github.com/pkg/errors"
)

type healthRepository struct {
	db *sqlx.DB
}

var _ repository.IHealthRepository = &healthRepository{}

// NewHealthRepository is constructor for healthRepository
func NewHealthRepository(db *sqlx.DB) repository.IHealthRepository {
	return &healthRepository{
		db: db,
	}
}

func (hr *healthRepository) Ping() error {
	query := `SELECT 1`
	_, err := hr.db.Exec(query)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
