// Package model is domain-model
package model

// User is User-Definition
type User struct {
	ID       int    `json:"id,omitempty" db:"id"`
	Name     string `json:"name,omitempty" db:"name"`
	Password string `json:"password,omitempty" db:"password"`
}
