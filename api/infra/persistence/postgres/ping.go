package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
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
