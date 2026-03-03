package user

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

type UserInput struct {
	Email        string `json:"email"`
	HashPassword string `json:"password"`
}

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (r *Repository) CreateUser(ctx context.Context, u UserInput) (string, error) {
	query := `

		INSERT INTO users (email, password_hash)
		VALUES ($1,$2)
		RETURNING id;
	`

	var id string

	err := r.db.QueryRow(
		ctx,
		query,
		u.Email,
		u.HashPassword,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	query := `
		SELECT id, email, password_hash
		FROM users
		WHERE email = $1;
	`

	var user User

	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		return User{}, nil
	}

	return user, nil

}
