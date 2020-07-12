// Package postgres is repository implementation package
package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
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

func (ur *userRepository) GetAll() ([]*model.User, error) {
	var uu []*model.User

	query := `SELECT
							id,
							name
						FROM
							public.user`

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
							public.user
						WHERE
							id = $1`

	if err := ur.db.Get(&user, query, userID); err != nil {
		return nil, errors.WithStack(err)
	}

	return &user, nil
}

func (ur *userRepository) GetByName(name string) ([]*model.User, error) {
	var users []*model.User

	query := `SELECT
							id,
							name
						FROM
							public.user
						WHERE
							name LIKE $1`

	if err := ur.db.Select(&users, query, "%"+name+"%"); err != nil {
		return nil, errors.WithStack(err)
	}

	return users, nil
}

func (ur *userRepository) Add(name string) (*model.User, error) {
	u := model.User{
		Name: name,
	}

	query := `INSERT INTO
							public.user (name)
						VALUES
							($1)`

	result, err := ur.db.Exec(query, name)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	i64, _ := result.LastInsertId()
	u.ID = int(i64)

	return &u, nil
}
