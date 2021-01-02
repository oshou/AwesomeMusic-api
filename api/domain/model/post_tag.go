// Package model is domain-model
package model

// PostTag represents N:N Relation for Post and Tag
type PostTag struct {
	PostID int `json:"post_id,omitempty" db:"post_id"`
	TagID  int `json:"tag_id,omitempty" db:"tag_id"`
}
