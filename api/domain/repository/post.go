// Package repository is domain-service
package repository

import "github.com/oshou/AwesomeMusic-api/api/domain/model"

// IPostRepository is Domain-access interface for Post
type IPostRepository interface {
	GetAll() ([]*model.Post, error)
	GetByID(postID int) (*model.Post, error)
	GetByTagID(tagID int) ([]*model.Post, error)
	GetByUserID(userID int) ([]*model.Post, error)
	Add(userID int, title, url, message string) (*model.Post, error)
	DeleteByID(postID int) error
}
