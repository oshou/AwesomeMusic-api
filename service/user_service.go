package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/entity"
)

type UserService struct{}

type User entity.User

func (us UserService) GetAll() ([]User, error) {
	var uu []User
	query := `SELECT
							id,
							name
						FROM
							user`
	conn := db.DBConn()
	if err := conn.Select(&uu, query); err != nil {
		return nil, err
	}
	return uu, nil
}

func (us UserService) GetById(user_id int) (User, error) {
	var u User
	query := `SELECT
							id,
							name
						FROM
							user
						WHERE
							id = ?`
	conn := db.DBConn()
	if err := conn.Get(&u, query, user_id); err != nil {
		return u, err
	}
	return u, nil
}

func (us UserService) GetByName(name string) ([]User, error) {
	var uu []User
	query := `SELECT
							id,
							name
						FROM
							user
						WHERE
							name LIKE ?`
	conn := db.DBConn()
	if err := conn.Select(&uu, query, "%"+name+"%"); err != nil {
		return uu, err
	}
	return uu, nil
}

func (us UserService) Add(name string) (User, error) {
	var u = User{
		Name: name,
	}
	query := `INSERT INTO
							user(name)
						VALUES
							(?)`
	conn := db.DBConn()
	result, err := conn.Exec(query, name)
	if err != nil {
		return u, err
	}
	i64, _ := result.LastInsertId()
	u.ID = int(i64)
	return u, nil
}
