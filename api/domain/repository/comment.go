//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
// Package repository is domain-service
package repository

import (
	"github.com/oshou/AwesomeMusic-api/api/domain/model"
)

// ICommentRepository is Domain-access interface for Comment
type ICommentRepository interface {
	List(postID int) ([]*model.Comment, error)
	GetByID(commentID int) (*model.Comment, error)
	Add(postID, userID int, comment string) (*model.Comment, error)
}
