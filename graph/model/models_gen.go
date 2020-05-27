// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type URL struct {
	ID          *int    `json:"ID"`
	Code        *string `json:"Code"`
	RedirectURL *string `json:"RedirectURL"`
	CreatedAt   *string `json:"CreatedAt"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
