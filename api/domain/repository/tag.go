//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/mock_$GOPACKAGE/$GOFILE
// Package repository is domain-service
package repository

import "github.com/oshou/AwesomeMusic-api/api/domain/model"

// ITagRepository is Domain-access interface for Tag & PostTag
type ITagRepository interface {
	List() ([]*model.Tag, error)
	ListByName(tagName string) ([]*model.Tag, error)
	ListByPostID(postID int) ([]*model.Tag, error)
	GetByID(tagID int) (*model.Tag, error)
	Add(tagName string) (*model.Tag, error)
	Attach(postID, tagID int) (*model.PostTag, error)
}
