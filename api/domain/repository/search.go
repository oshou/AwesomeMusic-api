//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
// Package repository is domain-service
package repository

import (
	"github.com/oshou/AwesomeMusic-api/api/domain/model"
)

// ISearchRepository is Domain-access interface for Post
type ISearchRepository interface {
	ListByTitle(q string) ([]*model.Post, error)
	ListByUserName(q string) ([]*model.Post, error)
	ListByTagName(q string) ([]*model.Post, error)
}
