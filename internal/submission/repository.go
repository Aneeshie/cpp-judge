package submission

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

type SubmissionInput struct {
	UserID    string
	ProblemId string
	Code      string
	Language  string
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
