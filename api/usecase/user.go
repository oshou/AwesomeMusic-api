//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
// Package usecase is application layer package
package usecase

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/log"
)

// IUserUsecase is usecase layer Interface for User
type IUserUsecase interface {
	Authenticate(username, password string) (*model.User, error)
	ListUsers() ([]*model.User, error)
	GetUserByID(userID int) (*model.User, error)
	AddUser(name, password string) (*model.User, error)
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
	if err != nil {
		e := errors.Cause(err)
		switch e.(type) {
		case repository.NoRowsError:
			return nil, NotFoundError{}
		default:
			log.Logger.Error("failed to get user by name", zap.Error(err))
			return nil, InternalServerError{}
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, UnauthorizedError{}
		}
		log.Logger.Error("failed to compare hash and password", zap.Error(err))
	}

	return user, nil
}

func (uu *userUsecase) ListUsers() ([]*model.User, error) {
	return uu.repo.List()
}

func (uu *userUsecase) GetUserByID(userID int) (*model.User, error) {
	return uu.repo.GetByID(userID)
}

func (uu *userUsecase) AddUser(name, password string) (*model.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Logger.Error("failed to generate password by bcrypt", zap.Error(err))
		return nil, InternalServerError{}
	}
	return uu.repo.Add(name, passwordHash)
}
