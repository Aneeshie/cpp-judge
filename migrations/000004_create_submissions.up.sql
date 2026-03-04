CREATE TYPE submission_verdict AS ENUM (
  'PENDING',
  'AC',
  'WA',
  'TLE',
  'MLE',
  'RE',
  'CE'
);

CREATE TABLE submissions (
    id UUID PRIMARY KEY gen_random_uuid(),

    user_id UUID NOT NULL,
    problem_id UUID NOT NULL,

    code TEXT NOT NULL,
    language TEXT NOT NULL,

    verdict submission_verdict NOT NULL DEFAULT 'PENDING',

    execution_time_ms INTEGER,
    memory_used_kb INTEGER,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_submission_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_submission_problem
        FOREIGN KEY (problem_id)
        REFERENCES problems(id)
        ON DELETE CASCADE
)
