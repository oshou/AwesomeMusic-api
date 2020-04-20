// Package service is application layer package
package service

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

// IUserService is service layer Interface for User
type IUserService interface {
	GetUsers() ([]*model.User, error)
	GetUserByID(userID int) (*model.User, error)
	AddUser(name string) (*model.User, error)
}

type userService struct {
	repo repository.IUserRepository
}

var _ IUserService = (*userService)(nil)

// NewUserService is IUserService constructor
func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{
		repo: repo,
	}
}

func (uu *userService) GetUsers() ([]*model.User, error) {
	return uu.repo.GetAll()
}

func (uu *userService) GetUserByID(userID int) (*model.User, error) {
	return uu.repo.GetByID(userID)
}

func (uu *userService) AddUser(name string) (*model.User, error) {
	return uu.repo.Add(name)
}
