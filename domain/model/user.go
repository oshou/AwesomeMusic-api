package model

type User struct {
	ID   int    `json:"id,omitempty" db:"id"`
	Name string `json:"name,omitempty" db:"name"`
}
