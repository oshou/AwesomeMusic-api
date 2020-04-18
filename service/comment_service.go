// Package service is application layer package
package service

import (
	"github.com/oshou/AwesomeMusic-api/domain/model"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
)

// ICommentService is service layer Interface for Comment
type ICommentService interface {
	GetComments(postID int) ([]*model.Comment, error)
	GetCommentByID(commentID int) (*model.Comment, error)
	AddComment(postID, userID int, comment string) (*model.Comment, error)
}

type commentService struct {
	repo repository.ICommentRepository
}

var _ ICommentService = (*commentService)(nil)

// NewCommentService is ICommentService constructor
func NewCommentService(repo repository.ICommentRepository) ICommentService {
	return &commentService{
		repo: repo,
	}
}

func (cu *commentService) GetComments(postID int) ([]*model.Comment, error) {
	return cu.repo.GetAll(postID)
}

func (cu *commentService) GetCommentByID(commentID int) (*model.Comment, error) {
	return cu.repo.GetByID(commentID)
}

func (cu *commentService) AddComment(postID, userID int, commentText string) (*model.Comment, error) {
	return cu.repo.Add(postID, userID, commentText)
}
