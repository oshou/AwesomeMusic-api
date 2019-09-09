package entity

type PostTag struct {
	PostID int `json:"post_id,omitempty"`
	TagID  int `json:"tag_id,omitempty"`
}
