//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../mock/mock_$GOPACKAGE/$GOFILE
// Package usecase is application layer package
package usecase

import (
	"github.com/oshou/AwesomeMusic-api/api/domain/model"
	"github.com/oshou/AwesomeMusic-api/api/domain/repository"
)

// ICommentUsecase is usecase layer Interface for Comment
type ICommentUsecase interface {
	ListComments(postID int) ([]*model.Comment, error)
	GetCommentByID(commentID int) (*model.Comment, error)
	AddComment(postID, userID int, comment string) (*model.Comment, error)
}

type commentUsecase struct {
	repo repository.ICommentRepository
}

var _ ICommentUsecase = &commentUsecase{}

// NewCommentUsecase is ICommentUsecase constructor
func NewCommentUsecase(repo repository.ICommentRepository) ICommentUsecase {
	return &commentUsecase{
		repo: repo,
	}
}

func (cu *commentUsecase) ListComments(postID int) ([]*model.Comment, error) {
	return cu.repo.List(postID)
}

func (cu *commentUsecase) GetCommentByID(commentID int) (*model.Comment, error) {
	return cu.repo.GetByID(commentID)
}

func (cu *commentUsecase) AddComment(postID, userID int, commentText string) (*model.Comment, error) {
	return cu.repo.Add(postID, userID, commentText)
}
