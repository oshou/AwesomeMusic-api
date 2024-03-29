//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/mock_$GOPACKAGE/$GOFILE
// Package usecase is application layer package
package usecase

import (
	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

// ITagUsecase is usecase layer Interface for Tag
type ITagUsecase interface {
	ListTags() ([]*model.Tag, error)
	ListTagsByPostID(postID int) ([]*model.Tag, error)
	GetTagByID(tagID int) (*model.Tag, error)
	AddTag(tagName string) (*model.Tag, error)
	AttachTag(postID, tagID int) (*model.PostTag, error)
}

type tagUsecase struct {
	repo repository.ITagRepository
}

var _ ITagUsecase = &tagUsecase{}

// NewTagUsecase is ITagUsecase constructor
func NewTagUsecase(repo repository.ITagRepository) ITagUsecase {
	return &tagUsecase{
		repo: repo,
	}
}

func (tu *tagUsecase) ListTags() ([]*model.Tag, error) {
	return tu.repo.List()
}

func (tu *tagUsecase) ListTagsByPostID(postID int) ([]*model.Tag, error) {
	return tu.repo.ListByPostID(postID)
}

func (tu *tagUsecase) GetTagByID(tagID int) (*model.Tag, error) {
	return tu.repo.GetByID(tagID)
}

func (tu *tagUsecase) AddTag(tagName string) (*model.Tag, error) {
	return tu.repo.Add(tagName)
}

func (tu *tagUsecase) AttachTag(postID, tagID int) (*model.PostTag, error) {
	return tu.repo.Attach(postID, tagID)
}
