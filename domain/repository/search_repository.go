// Package repository is domain-service
package repository

import "github.com/oshou/AwesomeMusic-api/domain/model"

type ISearchRepository interface {
	GetByTitle(q string) ([]*model.Post, error)
	GetByUserName(q string) ([]*model.Post, error)
	GetByTagName(q string) ([]*model.Post, error)
}
