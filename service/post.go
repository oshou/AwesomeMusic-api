// Package service is application layer package
package service

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

// IPostService is service layer Interface for Post
type IPostService interface {
	GetPosts() ([]*model.Post, error)
	GetPostByID(postID int) (*model.Post, error)
	GetPostsByTagID(tagID int) ([]*model.Post, error)
	GetPostsByUserID(userID int) ([]*model.Post, error)
	AddPost(userID int, title, url, message string) (*model.Post, error)
	DeletePostByID(postID int) error
}

type postService struct {
	repo repository.IPostRepository
}

var _ IPostService = &postService{}

// NewPostService is IPostService constructor
func NewPostService(repo repository.IPostRepository) IPostService {
	return &postService{
		repo: repo,
	}
}

func (pu *postService) GetPosts() ([]*model.Post, error) {
	return pu.repo.GetAll()
}

func (pu *postService) GetPostByID(postID int) (*model.Post, error) {
	return pu.repo.GetByID(postID)
}

func (pu *postService) GetPostsByTagID(tagID int) ([]*model.Post, error) {
	return pu.repo.GetByTagID(tagID)
}

func (pu *postService) GetPostsByUserID(userID int) ([]*model.Post, error) {
	return pu.repo.GetByUserID(userID)
}

func (pu *postService) AddPost(userID int, title, url, message string) (*model.Post, error) {
	return pu.repo.Add(userID, title, url, message)
}

func (pu *postService) DeletePostByID(postID int) error {
	return pu.repo.DeleteByID(postID)
}
