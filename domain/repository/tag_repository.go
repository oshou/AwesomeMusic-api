package repository

import "github.com/oshou/AwesomeMusic-api/domain/model"

type TagRepository interface {
	GetAll() (*model.Tag, error)
	GetByID(tagID int) (*model.Tag, error)
	GetByName(tagName string) ([]*model.Tag, error)
	GetByPostID(postID int) ([]*model.Tag, error)
	Add(tagName string) (*model.Tag, error)
	Attach(postID, tagID int) (*model.PostTag, error)
}
