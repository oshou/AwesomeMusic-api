package datastore

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type userRepository struct {
	DB *sqlx.DB
}

var _ repository.IUserRepository = (*userRepository)(nil)

func NewUserRepository(db *sqlx.DB) repository.IUserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) GetAll() ([]*model.User, error) {
	var uu []*model.User

	query := `SELECT
							id,
							name
						FROM
							user`

	if err := ur.DB.Select(&uu, query); err != nil {
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

	if err := ur.DB.Get(u, query, userID); err != nil {
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

	if err := ur.DB.Select(&uu, query, "%"+name+"%"); err != nil {
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

	result, err := ur.DB.Exec(query, name)

	if err != nil {
		return u, err
	}

	i64, _ := result.LastInsertId()
	u.ID = int(i64)

	return u, nil
}
