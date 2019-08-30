package entity

type Comment struct {
	Id      int64  `json:"id,omitempty"`
	PostId  int64  `json:"post_id,omitempty"`
	comment string `json:"comment,omitempty"`
}
