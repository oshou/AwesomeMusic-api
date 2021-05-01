//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
// Package repository is domain-service
package repository

import (
	"github.com/oshou/AwesomeMusic-api/api/domain/model"
)

// ISearchRepository is Domain-access interface for Post
type ISearchRepository interface {
	GetByTitle(q string) ([]*model.Post, error)
	GetByUserName(q string) ([]*model.Post, error)
	GetByTagName(q string) ([]*model.Post, error)
}
