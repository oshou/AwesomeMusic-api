//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/mock_$GOPACKAGE/$GOFILE
// Package repository is domain-service
package repository

import "github.com/oshou/AwesomeMusic-api/api/domain/model"

// IPostRepository is Domain-access interface for Post
type IPostRepository interface {
	List() ([]*model.Post, error)
	ListByTagID(tagID int) ([]*model.Post, error)
	ListByUserID(userID int) ([]*model.Post, error)
	GetByID(postID int) (*model.Post, error)
	Add(userID int, title, url, message string) (*model.Post, error)
	DeleteByID(postID int) error
}
