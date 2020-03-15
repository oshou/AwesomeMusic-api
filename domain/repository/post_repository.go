package repository

import "github.com/oshou/AwesomeMusic-api/domain/model"

type PostRepository interface {
	GetAll() ([]model.Post, error)
	GetByID(postID int) (model.Post, error)
	GetByTagID(tagID int) ([]model.Post, error)
	GetByUserID(userID int) ([]model.Post, error)
	Add(userID int, title, url, message string) (model.Post, error)
	DeleteByID(postID int) error
}
