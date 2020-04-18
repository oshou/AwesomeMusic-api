// Package repository is domain-service
package repository

import "github.com/oshou/AwesomeMusic-api/domain/model"

// ICommentRepository is Domain-access interface for Comment
type ICommentRepository interface {
	GetAll(postID int) ([]*model.Comment, error)
	GetByID(commentID int) (*model.Comment, error)
	Add(postID, userID int, comment string) (*model.Comment, error)
}
