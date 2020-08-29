// Package usecase is application layer package
package usecase

import (
	"database/sql"

	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/log"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// IUserUsecase is usecase layer Interface for User
type IUserUsecase interface {
	GetUsers() ([]*model.User, error)
	GetUserByID(userID int) (*model.User, error)
	AddUser(name string) (*model.User, error)
	Authenticate(username, password string) (*model.User, error)
}

type userUsecase struct {
	repo repository.IUserRepository
}

var _ IUserUsecase = &userUsecase{}

// NewUserUsecase is IUserUsecase constructor
func NewUserUsecase(repo repository.IUserRepository) IUserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (uu *userUsecase) Authenticate(username, password string) (*model.User, error) {
	user, err := uu.repo.GetByName(username)
	if errors.Cause(err) == sql.ErrNoRows {
		return nil, UnauthorizedError{}
	}
	if err != nil {
		log.Logger.Error("failed to get user by name", zap.Error(err))
		return nil, InternalServerError{}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, UnauthorizedError{}
		}
		log.Logger.Error("failed to compare hash and password", zap.Error(err))
	}
	return user, nil
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
