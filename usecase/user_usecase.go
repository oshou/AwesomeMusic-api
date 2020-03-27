package service

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type UserService interface {
	GetUsers() ([]*model.User, error)
	GetUserByID() (*model.User, error)
	AddUser() (*model.User, error)
}

func NewUserService(repo repository.UserRepository, pre presenter.UserPresenter) UserServiceInterface {
	return &userService{
		UserRepository: repo,
		UserPresenter:  pre,
	}
}

type UserService struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

func (us *userService) GetUsers(uu []*model.User) ([]*model.User, error) {
	return userService.UserRepository.GetUsers(u)
}

func (us *userService) GetUserByID(userID int, u *model.User) (*model.User, error) {
	return UserService.UserRepository.GetUserByID(userID)
}

func (us *userService) AddUser(name string, u *model.User) (*model.User, error) {
	return userService.UserRepository.Add(name)
}
