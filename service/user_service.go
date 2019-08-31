package service

import (
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/entity"
)

type UserService struct{}

type User entity.User

func (us UserService) GetAll() ([]User, error) {

	var u []User

	stmt := db.GetDBConn()
	stmt = stmt.Table("user")
	stmt = stmt.Select("id,name")
	if err := stmt.Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (us UserService) GetById(user_id int) (User, error) {

	var u User

	stmt := db.GetDBConn()
	stmt = stmt.Table("user")
	stmt = stmt.Select("id,name")
	stmt = stmt.Where("id = ?", user_id)
	if err := stmt.First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (us UserService) Add(name string) (User, error) {

	var u User
	u.Name = name

	stmt := db.GetDBConn()
	stmt = stmt.Table("user")
	if err := stmt.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}
