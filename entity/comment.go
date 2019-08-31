package entity

type Comment struct {
	ID      int    `json:"id,omitempty"`
	UserID  int    `json:"user_id,omitempty"`
	PostID  int    `json:"post_id,omitempty"`
	Comment string `json:"comment,omitempty"`
}
