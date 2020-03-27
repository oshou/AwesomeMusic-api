package controller

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/usecase/service"
)

type UserControllerInterface interface {
	GetUsers() ([]*model.User,error)
	GetUserByID() (*model.User, error)
	AddUser() (*model.User, error)
}

func NewUserController(us service.UserServiceInterface) UserControllerInterface {
	return UserController(us)
}

type UserController struct{
	userService: service.UserService
}

func (uc *UserController)GetUsers() ([]*model.User,error){
	return uc.UserService.GetUsers()
}
