CREATE TABLE problems (
  id UUID PRIMARY KEY gen_random_uuid(),
  title CITEXT NOT NULL UNIQUE,
  description TEXT NOT NULL,
  time_limit_ms INTEGER NOT NULL,
  memory_limit_ms INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
)

