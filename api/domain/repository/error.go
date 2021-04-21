package repository

import (
	"database/sql"
)

type NoRowsError struct{}

func (NoRowsError) Error() string {
	return sql.ErrNoRows.Error()
}
