// Package service is application layer package
package service

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

// ISearchService is service layer Interface
type ISearchService interface {
	GetPostsByTitle(q string) ([]*model.Post, error)
	GetPostsByUserName(q string) ([]*model.Post, error)
	GetPostsByTagName(q string) ([]*model.Post, error)
}

type searchService struct {
	repo repository.ISearchRepository
}

var _ ISearchService = (*searchService)(nil)

// NewSearchService is ISearchService constructor
func NewSearchService(repo repository.ISearchRepository) ISearchService {
	return &searchService{
		repo: repo,
	}
}

func (su *searchService) GetPostsByTitle(q string) ([]*model.Post, error) {
	return su.repo.GetByTitle(q)
}

func (su *searchService) GetPostsByUserName(q string) ([]*model.Post, error) {
	return su.repo.GetByUserName(q)
}

func (su *searchService) GetPostsByTagName(q string) ([]*model.Post, error) {
	return su.repo.GetByTagName(q)
}
