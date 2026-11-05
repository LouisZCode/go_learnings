-- Runs automatically on first start of the Postgres container
-- (mounted into /docker-entrypoint-initdb.d by docker-compose.yml).

CREATE TABLE IF NOT EXISTS items (
    id         BIGINT      GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title      TEXT        NOT NULL,
    done       BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
