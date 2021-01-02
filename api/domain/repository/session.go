// Package repository is domain-service
package repository

import (
	"github.com/oshou/AwesomeMusic-api/api/domain/model"
)

// ISearchRepository is Domain-access interface for Post
type ISessionRepository interface {
	Get(id string) (*model.Session, error)
	Set(*model.Session) error
}
