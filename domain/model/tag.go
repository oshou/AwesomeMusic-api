// Package model is domain-model
package model

type Tag struct {
	ID   int    `json:"id,omitempty" db:"id"`
	Name string `json:"name,omitempty" db:"name"`
}
