package entity

type User struct {
	ID   int    `json:"id,imitempty"`
	Name string `json:"name,omitempty"`
}
