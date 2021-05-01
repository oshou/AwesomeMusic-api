//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
// Package usecase is application layer package
package usecase

import (
	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

// ISearchUsecase is usecase layer Interface
type ISearchUsecase interface {
	ListPostsByTitle(q string) ([]*model.Post, error)
	ListPostsByUserName(q string) ([]*model.Post, error)
	ListPostsByTagName(q string) ([]*model.Post, error)
}

type searchUsecase struct {
	repo repository.ISearchRepository
}

var _ ISearchUsecase = &searchUsecase{}

// NewSearchUsecase is ISearchUsecase constructor
func NewSearchUsecase(repo repository.ISearchRepository) ISearchUsecase {
	return &searchUsecase{
		repo: repo,
	}
}

func (su *searchUsecase) ListPostsByTitle(q string) ([]*model.Post, error) {
	return su.repo.ListByTitle(q)
}

func (su *searchUsecase) ListPostsByUserName(q string) ([]*model.Post, error) {
	return su.repo.ListByUserName(q)
}

func (su *searchUsecase) ListPostsByTagName(q string) ([]*model.Post, error) {
	return su.repo.ListByTagName(q)
}
