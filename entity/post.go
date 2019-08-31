package entity

type Post struct {
	ID      int    `json:"id,omitempty"`
	UserID  int    `json:"user_id,omitempty"`
	URL     string `json:"url,omitempty"`
	Message string `json:"message,omitempty"`
}
