package datastore

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/model"
)

type UserStore struct {
	DB *sqlx.DB
}

func (us *UserStore) GetAll() ([]model.User, error) {
	var uu []model.User

	query := `SELECT
							id,
							name
						FROM
							user`

	if err := us.DB.Select(&uu, query); err != nil {
		return nil, err
	}

	return uu, nil
}

func (us *UserStore) GetByID(userID int) (model.User, error) {
	var u model.User

	query := `SELECT
							id,
							name
						FROM
							user
						WHERE
							id = ?`

	if err := us.DB.Get(&u, query, userID); err != nil {
		return u, err
	}

	return u, nil
}

func (us *UserStore) GetByName(name string) ([]model.User, error) {
	var uu []model.User

	query := `SELECT
							id,
							name
						FROM
							user
						WHERE
							name LIKE ?`

	if err := us.DB.Select(&uu, query, "%"+name+"%"); err != nil {
		return uu, err
	}

	return uu, nil
}

func (us *UserStore) Add(name string) (model.User, error) {
	var u = model.User{
		Name: name,
	}

	query := `INSERT INTO
							user(name)
						VALUES
							(?)`

	result, err := us.DB.Exec(query, name)

	if err != nil {
		return u, err
	}

	i64, _ := result.LastInsertId()
	u.ID = int(i64)

	return u, nil
}
