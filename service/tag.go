// Package service is application layer package
package service

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

// ITagService is service layer Interface for Tag
type ITagService interface {
	GetTags() ([]*model.Tag, error)
	GetTagByID(tagID int) (*model.Tag, error)
	GetTagsByPostID(postID int) ([]*model.Tag, error)
	AddTag(tagName string) (*model.Tag, error)
	AttachTag(postID, tagID int) (*model.PostTag, error)
}

type tagService struct {
	repo repository.ITagRepository
}

var _ ITagService = (*tagService)(nil)

// NewTagService is ITagService constructor
func NewTagService(repo repository.ITagRepository) ITagService {
	return &tagService{
		repo: repo,
	}
}

func (tu *tagService) GetTags() ([]*model.Tag, error) {
	return tu.repo.GetAll()
}

func (tu *tagService) GetTagByID(tagID int) (*model.Tag, error) {
	return tu.repo.GetByID(tagID)
}

func (tu *tagService) GetTagsByPostID(postID int) ([]*model.Tag, error) {
	return tu.repo.GetByPostID(postID)
}

func (tu *tagService) AddTag(tagName string) (*model.Tag, error) {
	return tu.repo.Add(tagName)
}

func (tu *tagService) AttachTag(postID, tagID int) (*model.PostTag, error) {
	return tu.repo.Attach(postID, tagID)
}
