package users

import "context"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Position string `json:"position"`
}

type Repository interface {
	GetUser(ctx context.Context, id int) (User, error)
}
