package entity

type User struct {
	Id   int64  `json:"id,imitempty"`
	Name string `json:"name,omitempty"`
}
