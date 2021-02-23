package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/keepcalmist/hospital/pkg/users"
)

type userRepository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) users.Repository {
	return &userRepository{db: db}
}

func (repo *userRepository) GetUser(ctx context.Context, id int) (users.User, error) {
	query := `SELECT * FROM users WHERE  id = ?`
	user := users.User{}
	row := repo.db.QueryRowContext(ctx, query, id)

	if row.Err() != nil {
		return users.User{}, row.Err()
	}
	if err := row.Scan(&user); err != nil {
		return users.User{}, err
	}
	return user, nil
}
