// Package repository is domain-service
package repository

import "github.com/oshou/AwesomeMusic-api/domain/model"

type ICommentRepository interface {
	GetAll(postID int) ([]*model.Comment, error)
	GetByID(commentID int) (*model.Comment, error)
	Add(postID, userID int, comment string) (*model.Comment, error)
}
