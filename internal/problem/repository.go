package problem

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

type Problem struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	TimeLimitMs   int    `json:"time_limit_ms"`
	MemoryLimitMb int    `json:"memory_limit_mb"`
}

func (r *Repository) CreateProblem(ctx context.Context, p Problem) (Problem, error) {
	query := `
		INSERT INTO problems (title, description, time_limit_ms, memory_limit_mb)
		VALUES ($1,$2,$3,$4)
		RETURNING id;
	`

	err := r.db.QueryRow(
		ctx,
		query,
		p.Title,
		p.Description,
		p.TimeLimitMs,
		p.MemoryLimitMb,
	).Scan(&p.ID)

	if err != nil {
		return Problem{}, err
	}

	return p, nil

}

func (r *Repository) GetProblems(ctx context.Context) ([]Problem, error) {
	query := `
		SELECT id, title, description, time_limit_ms, memory_limit_mb, created_at
		FROM problems
		WHERE deleted_at IS NULL;
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var problems []Problem

	for rows.Next() {
		var p Problem
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.TimeLimitMs,
			&p.MemoryLimitMb,
		)
		if err != nil {
			return nil, err
		}
		problems = append(problems, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return problems, nil
}
