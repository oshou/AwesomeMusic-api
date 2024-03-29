//go:generate mockgen -source=$GOFILE -package=mock_$GOPACKAGE -destination=../../mock/mock_$GOPACKAGE/$GOFILE
// Package repository is domain-service
package repository

import "github.com/oshou/AwesomeMusic-api/api/domain/model"

// IUserRepository is Domain-access interface for User
type IUserRepository interface {
	List() ([]*model.User, error)
	GetByID(userID int) (*model.User, error)
	GetByName(name string) (*model.User, error)
	Add(name string, passwordHash []byte) (*model.User, error)
}
