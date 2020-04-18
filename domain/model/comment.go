// Package model is domain-model
package model

// Comment is comment to AwesomeMusic posts
type Comment struct {
	ID      int    `json:"id,omitempty" db:"id"`
	UserID  int    `json:"user_id,omitempty" db:"user_id"`
	PostID  int    `json:"post_id,omitempty" db:"post_id"`
	Comment string `json:"comment,omitempty" db:"comment"`
}
