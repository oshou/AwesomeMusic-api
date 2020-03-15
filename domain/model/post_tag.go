package model

type PostTag struct {
	PostID int `json:"post_id,omitempty" db:"post_id"`
	TagID  int `json:"tag_id,omitempty" db:"tag_id"`
}
