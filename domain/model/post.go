// Package model is domain-model
package model

// Post is AwesomeMusic Post
type Post struct {
	ID      int    `json:"id,omitempty" db:"id"`
	UserID  int    `json:"user_id,omitempty" db:"user_id"`
	Title   string `json:"title,omitempty" db:"title"`
	URL     string `json:"url,omitempty" db:"url"`
	Message string `json:"message,omitempty" db:"message"`
}
