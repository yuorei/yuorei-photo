// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserPayload struct {
	User *User `json:"user"`
}

type Mutation struct {
}

type Query struct {
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}