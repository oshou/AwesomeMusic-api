//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
// Package usecase is application layer package
package usecase

import (
	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

// IPostUsecase is usecase layer Interface for Post
type IPostUsecase interface {
	ListPosts() ([]*model.Post, error)
	ListPostsByTagID(tagID int) ([]*model.Post, error)
	ListPostsByUserID(userID int) ([]*model.Post, error)
	GetPostByID(postID int) (*model.Post, error)
	AddPost(userID int, title, url, message string) (*model.Post, error)
	DeletePostByID(postID int) error
}

type postUsecase struct {
	repo repository.IPostRepository
}

var _ IPostUsecase = &postUsecase{}

// NewPostUsecase is IPostUsecase constructor
func NewPostUsecase(repo repository.IPostRepository) IPostUsecase {
	return &postUsecase{
		repo: repo,
	}
}

func (pu *postUsecase) ListPosts() ([]*model.Post, error) {
	return pu.repo.List()
}

func (pu *postUsecase) GetPostByID(postID int) (*model.Post, error) {
	return pu.repo.GetByID(postID)
}

func (pu *postUsecase) ListPostsByTagID(tagID int) ([]*model.Post, error) {
	return pu.repo.GetByTagID(tagID)
}

func (pu *postUsecase) ListPostsByUserID(userID int) ([]*model.Post, error) {
	return pu.repo.GetByUserID(userID)
}

func (pu *postUsecase) AddPost(userID int, title, url, message string) (*model.Post, error) {
	return pu.repo.Add(userID, title, url, message)
}

func (pu *postUsecase) DeletePostByID(postID int) error {
	return pu.repo.DeleteByID(postID)
}
