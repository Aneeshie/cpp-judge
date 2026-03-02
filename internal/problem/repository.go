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

type SubmissionInput struct {
	UserID    string
	ProblemId string
	Code      string
	Language  string
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

func (r *Repository) MakeSubmission(ctx context.Context, input SubmissionInput) (string, error) {
	tx, err := r.db.Begin(ctx)

	if err != nil {
		return "", err
	}

	defer tx.Rollback(ctx)

	var submissionId string

	query := `
		INSERT INTO submissions (user_id, problem_id, code, language,verdict)
		VALUES ($1,$2,$3,$4, 'PENDING')
		RETURNING id;
	`

	err = tx.QueryRow(
		ctx,
		query,
		input.UserID,
		input.ProblemId,
		input.Code,
		input.Language,
	).Scan(&submissionId)

	if err != nil {
		return "", err
	}

	insertJobQuery := `
		INSERT INTO submission_jobs (submission_id, status)
		VALUES ($1,'PENDING')
	`
	_, err = tx.Exec(ctx, insertJobQuery, submissionId)

	if err != nil {
		return "", err
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}

	return submissionId, nil
}
