// Package postgres is repository implementation package
package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

type userRepository struct {
	db *sqlx.DB
}

var _ repository.IUserRepository = &userRepository{}

// NewUserRepository is constructor for userRepository
func NewUserRepository(db *sqlx.DB) repository.IUserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) List() ([]*model.User, error) {
	var uu []*model.User

	query := `SELECT
							id,
							name
						FROM
							public.users`

	if err := ur.db.Select(&uu, query); err != nil {
		return nil, errors.WithStack(err)
	}

	return uu, nil
}

func (ur *userRepository) GetByID(userID int) (*model.User, error) {
	var user model.User

	query := `SELECT
							id,
							name
						FROM
							public.users
						WHERE
							id = $1`

	if err := ur.db.Get(&user, query, userID); err != nil {
		return nil, errors.WithStack(err)
	}

	return &user, nil
}

func (ur *userRepository) GetByName(name string) (*model.User, error) {
	var user model.User

	query := `SELECT
							id,
							name,
							password_hash
						FROM
							public.users
						WHERE
							name LIKE $1`

	if err := ur.db.Get(&user, query, "%"+name+"%"); err != nil {
		return nil, errors.WithStack(err)
	}

	return &user, nil
}

func (ur *userRepository) Add(name string, passwordHash []byte) (*model.User, error) {
	u := model.User{
		Name: name,
	}

	query := `INSERT INTO
							public.users(name,password_hash)
						VALUES
							($1, $2)
						RETURNING
							id`

	err := ur.db.QueryRow(query, name, passwordHash).Scan(&u.ID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &u, nil
}
