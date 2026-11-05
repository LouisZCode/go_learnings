# listapp — CRUD list app (Go + Vite/React + Postgres)

Small reference implementation of the capstone: create, update, and delete items in a list.

## Architecture

    Browser ── /api/* ──▶ Go server (:8080) ── SQL ──▶ Postgres (:5432)
       ▲                     │
       └── static files ─────┘   (the built frontend is served by the same Go binary)

- `main.go` — HTTP server, routes, JSON handling (standard library only, Go 1.22+ routing)
- `internal/store/store.go` — all SQL lives here (pgx v5 driver)
- `db/init.sql` — table schema
- `web/` — Vite + React frontend (`/api` requests are proxied to the Go server in dev)

API: `GET /api/items` · `POST /api/items` · `PUT /api/items/{id}` · `DELETE /api/items/{id}`

## Prerequisites

- Go 1.22+ (developed on 1.27)
- Node 20+
- A local PostgreSQL (any recent major version)

## One-time database setup

Creates the `listapp` role (password `listapp`), the database, and the `items` table.

**macOS — PostgreSQL 18 from the official installer:**

```bash
/Library/PostgreSQL/18/bin/psql -h localhost -U postgres \
  -c "CREATE ROLE listapp LOGIN PASSWORD 'listapp';" \
  -c "CREATE DATABASE listapp OWNER listapp;"

PGPASSWORD=listapp /Library/PostgreSQL/18/bin/psql -h localhost -U listapp -d listapp -f db/init.sql
```

**WSL / Linux (team environment):**

```bash
sudo -u postgres psql \
  -c "CREATE ROLE listapp LOGIN PASSWORD 'listapp';" \
  -c "CREATE DATABASE listapp OWNER listapp;"
sudo -u postgres psql -d listapp -f db/init.sql
```

## Run it

```bash
go run .                      # terminal 1 — serves API + built frontend on :8080
cd web
npm install                   # once
npm run dev                   # terminal 2 — Vite dev server on :5173, proxies /api to :8080
```

- Dev mode: open **http://localhost:5173**
- Production-style mode (single Go process): `npm run build` in `web/`, then open **http://localhost:8080**

## Configuration

Environment variables (both optional):

| Variable       | Default                                                          |
|----------------|------------------------------------------------------------------|
| `PORT`         | `8080`                                                           |
| `DATABASE_URL` | `postgres://listapp:listapp@localhost:5432/listapp?sslmode=disable` |
