// Package usecase is application layer package
package usecase

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

// ICommentUsecase is usecase layer Interface for Comment
type ICommentUsecase interface {
	GetComments(postID int) ([]*model.Comment, error)
	GetCommentByID(commentID int) (*model.Comment, error)
	AddComment(postID, userID int, comment string) (*model.Comment, error)
}

type commentUsecase struct {
	repo repository.ICommentRepository
}

var _ ICommentUsecase = (*commentUsecase)(nil)

// NewCommentUsecase is ICommentUsecase constructor
func NewCommentUsecase(repo repository.ICommentRepository) ICommentUsecase {
	return &commentUsecase{
		repo: repo,
	}
}

func (cu *commentUsecase) GetComments(postID int) ([]*model.Comment, error) {
	return cu.repo.GetAll(postID)
}

func (cu *commentUsecase) GetCommentByID(commentID int) (*model.Comment, error) {
	return cu.repo.GetByID(commentID)
}

func (cu *commentUsecase) AddComment(postID, userID int, commentText string) (*model.Comment, error) {
	return cu.repo.Add(postID, userID, commentText)
}
