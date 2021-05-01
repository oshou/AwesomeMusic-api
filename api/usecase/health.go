//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import "github.com/oshou/AwesomeMusic-api/api/domain/repository"

type IHealthUsecase interface {
	GetHealth() error
}

type healthUsecase struct {
	repo repository.IHealthRepository
}

var _ IHealthUsecase = &healthUsecase{}

func NewHealthUsecase(repo repository.IHealthRepository) IHealthUsecase {
	return &healthUsecase{
		repo: repo,
	}
}

func (hu *healthUsecase) GetHealth() error {
	return hu.repo.Ping()
}
