// Package usecase is application layer package
package usecase

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

// ITagUsecase is usecase layer Interface for Tag
type ITagUsecase interface {
	GetTags() ([]*model.Tag, error)
	GetTagByID(tagID int) (*model.Tag, error)
	GetTagsByPostID(postID int) ([]*model.Tag, error)
	AddTag(tagName string) (*model.Tag, error)
	AttachTag(postID, tagID int) (*model.PostTag, error)
}

type tagUsecase struct {
	repo repository.ITagRepository
}

var _ ITagUsecase = (*tagUsecase)(nil)

// NewTagUsecase is ITagUsecase constructor
func NewTagUsecase(repo repository.ITagRepository) ITagUsecase {
	return &tagUsecase{
		repo: repo,
	}
}

func (tu *tagUsecase) GetTags() ([]*model.Tag, error) {
	return tu.repo.GetAll()
}

func (tu *tagUsecase) GetTagByID(tagID int) (*model.Tag, error) {
	return tu.repo.GetByID(tagID)
}

func (tu *tagUsecase) GetTagsByPostID(postID int) ([]*model.Tag, error) {
	return tu.repo.GetByPostID(postID)
}

func (tu *tagUsecase) AddTag(tagName string) (*model.Tag, error) {
	return tu.repo.Add(tagName)
}

func (tu *tagUsecase) AttachTag(postID, tagID int) (*model.PostTag, error) {
	return tu.repo.Attach(postID, tagID)
}
