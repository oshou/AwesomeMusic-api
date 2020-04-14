package usecase

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

type IPostUsecase interface {
	GetPosts() ([]*model.Post, error)
	GetPostByID(postID int) (*model.Post, error)
	GetPostsByTagID(tagID int) ([]*model.Post, error)
	GetPostsByUserID(userID int) ([]*model.Post, error)
	AddPost(userID int, title, url, message string) (*model.Post, error)
	DeletePostByID(postID int) error
}

type postUsecase struct {
	repo repository.IPostRepository
}

var _ IPostUsecase = (*postUsecase)(nil)

func NewPostUsecase(repo repository.IPostRepository) IPostUsecase {
	return &postUsecase{
		repo: repo,
	}
}

func (pu *postUsecase) GetPosts() ([]*model.Post, error) {
	return pu.repo.GetAll()
}

func (pu *postUsecase) GetPostByID(postID int) (*model.Post, error) {
	return pu.repo.GetByID(postID)
}

func (pu *postUsecase) GetPostsByTagID(tagID int) ([]*model.Post, error) {
	return pu.repo.GetByTagID(tagID)
}

func (pu *postUsecase) GetPostsByUserID(userID int) ([]*model.Post, error) {
	return pu.repo.GetByUserID(userID)
}

func (pu *postUsecase) AddPost(userID int, title, url, message string) (*model.Post, error) {
	return pu.repo.Add(userID, title, url, message)
}

func (pu *postUsecase) DeletePostByID(postID int) error {
	return pu.repo.DeleteByID(postID)
}
