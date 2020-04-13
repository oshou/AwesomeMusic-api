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

type userService struct {
	repo repository.UserRepository
	pre  presenter.UserPresenter
}

var _ UserService = &userService{}

func NewUserService(repo repository.UserRepository, pre presenter.UserPresenter) UserService {
	return &userService{
		repo: repo,
		pre:  pre,
	}
}

func (us *userService) GetUsers(uu []*model.User) ([]*model.User, error) {
	users := us.repo.GetAll()
	return us.pre.ResponseUsers(users)
}

func (us *userService) GetUserByID(userID int, u *model.User) (*model.User, error) {
	user := us.repo.GetByID()
	return us.pre.ResponseUser(us)
}

func (us *userService) AddUser(name string, u *model.User) (*model.User, error) {
	return us.pre.ResponseUsers(user)
}
