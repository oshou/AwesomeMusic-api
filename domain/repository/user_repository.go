package repository

import "github.com/oshou/AwesomeMusic-api/domain/model"

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetByID(userID int) (model.User, error)
	GetByName(name string) ([]model.User, error)
	Add(name string) (model.User, error)
}
