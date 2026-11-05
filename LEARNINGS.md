# LEARNINGS

Discoveries and decisions worth keeping across sessions. Append as modules complete.

## 2026-11-04 — Restart with a target

- Luis restarted the curriculum: goal is to learn Go completely for his new job, starting with a fast review of what is already covered (modules 0–1).
- Capstone defined: small CRUD list app — Go web server + Vite/React frontend + Postgres — to be presented at work soon. Every module now aims at it; recorded in `AGENTS.md` (renamed from `CLAUDE.md` this session).
- Environment: Luis develops on macOS (arm64, go1.27.0); the work team uses WSL on Windows. Everything stays portable; Postgres will run via Docker.
- Git verified: remote `github.com/LouisZCode/go_learnings` reachable, push works.
- Notion MCP was not connected this session — the 🧿 GO notes page could not be reviewed. Revisit next session.
- Learning sources locked (2026-11-05): the two course videos are downloaded to `media/` with `.transcript.md` sidecars generated from their embedded YouTube captions (91.7k and 59.6k words). `media/` is gitignored — videos and transcripts live locally only.
- Capstone reference implementation built and committed at `listapp/` (own Go module; moved into the repo root on Luis's request). Single Go dependency: pgx v5; frontend is plain Vite + React. No Docker anywhere — local Postgres only. The Mac already ran PostgreSQL 18 (official installer, password auth everywhere, `postgres` OS user exists → `sudo -u postgres psql` is the admin path).
- Backend walkthrough started (Lesson B1: request → router → handler → response). Follow-up track: `listapp-evolution/` — the app rebuilt in steps Luis can read and run; step 0 (`00_basic_http`, std-lib only, in-memory list, port 8081) built, tested end-to-end, and left running.
- Step 02 (`02_postgres`, port 8082, own go.mod + pgx v5) done and verified: same API as step 0, storage swapped for pgxpool + SQL. Shares the `listapp` database, so rows added in the app's UI show up in step 02 and vice versa. Lesson B2 (how Go talks to Postgres) delivered before building it, per Luis's explain-first request.
