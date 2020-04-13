package usecase

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type UserUsecase interface {
	GetUsers() ([]*model.User, error)
	GetUserByID() (*model.User, error)
	AddUser() (*model.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

var _ UserUsecase = &userUsecase{}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (uu *userUsecase) GetUsers(uu []*model.User) ([]*model.User, error) {
	return uu.repo.GetAll()
}

func (uu *userUsecase) GetUserByID(userID int, u *model.User) (*model.User, error) {
	return uu.repo.GetByID()
}

func (uu *userUsecase) AddUser(name string, u *model.User) (*model.User, error) {
	return uu.repo.Add(name)
}
