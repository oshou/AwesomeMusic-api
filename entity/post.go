package entity

type Post struct {
	Id      int64  `json:"id,omitempty"`
	UserId  int64  `json:"user_id,omitempty`
	Url     string `json:"url,omitempty"`
	Message string `json:"message,omitempty"`
}
