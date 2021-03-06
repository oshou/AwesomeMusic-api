// Package usecase is application layer package
package usecase

import (
	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

// ISearchUsecase is usecase layer Interface
type ISearchUsecase interface {
	GetPostsByTitle(q string) ([]*model.Post, error)
	GetPostsByUserName(q string) ([]*model.Post, error)
	GetPostsByTagName(q string) ([]*model.Post, error)
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

func (su *searchUsecase) GetPostsByTitle(q string) ([]*model.Post, error) {
	return su.repo.GetByTitle(q)
}

func (su *searchUsecase) GetPostsByUserName(q string) ([]*model.Post, error) {
	return su.repo.GetByUserName(q)
}

func (su *searchUsecase) GetPostsByTagName(q string) ([]*model.Post, error) {
	return su.repo.GetByTagName(q)
}
