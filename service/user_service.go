package user

import (
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/entity"
)

type Service struct{}

type User entity.User

func (s Service) GetUser() ([]User, error) {

	var u []User

	stmt := db.GetDBConn()
	stmt = stmt.Table("user")
	stmt = stmt.Select("id,name")
	if err := stmt.Find(&u).Error; err != nil {
		return nil,err
	}
	return u,nil
}

func (s Service) AddUser() (User, error) {

	var u User

	if err := c.BindJSON(&u); err != nil{
		return u,err
	}

	stmt := db.GetDBConn()
	if err :=stmt.Create(&u).Error; err !=nil{
		return u,err
	}
	return u,nil
}
