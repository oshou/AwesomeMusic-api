// Package mysql is repository implementation package
package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type userRepository struct {
	db *sqlx.DB
}

var _ repository.IUserRepository = (*userRepository)(nil)

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
							user`

	if err := ur.db.Select(&uu, query); err != nil {
		return nil, err
	}

	return uu, nil
}

func (ur *userRepository) GetByID(userID int) (*model.User, error) {
	u := &model.User{}

	query := `SELECT
							id,
							name
						FROM
							user
						WHERE
							id = ?`

	if err := ur.db.Get(u, query, userID); err != nil {
		return u, err
	}

	return u, nil
}

func (ur *userRepository) GetByName(name string) ([]*model.User, error) {
	var uu []*model.User

	query := `SELECT
							id,
							name
						FROM
							user
						WHERE
							name LIKE ?`

	if err := ur.db.Select(&uu, query, "%"+name+"%"); err != nil {
		return uu, err
	}

	return uu, nil
}

func (ur *userRepository) Add(name string) (*model.User, error) {
	u := &model.User{
		Name: name,
	}
	query := `INSERT INTO
							user(name)
						VALUES
							(?)`

	result, err := ur.db.Exec(query, name)

	if err != nil {
		return u, err
	}

	i64, _ := result.LastInsertId()
	u.ID = int(i64)

	return u, nil
}
