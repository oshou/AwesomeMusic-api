package usecase

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type IUserUsecase interface {
	GetUsers() ([]*model.User, error)
	GetUserByID(userID int) (*model.User, error)
	AddUser(name string) (*model.User, error)
}

type userUsecase struct {
	repo repository.IUserRepository
}

var _ IUserUsecase = (*userUsecase)(nil)

func NewUserUsecase(repo repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (uu *userUsecase) GetUsers() ([]*model.User, error) {
	return uu.repo.GetAll()
}

func (uu *userUsecase) GetUserByID(userID int) (*model.User, error) {
	return uu.repo.GetByID(userID)
}

func (uu *userUsecase) AddUser(name string) (*model.User, error) {
	return uu.repo.Add(name)
}
