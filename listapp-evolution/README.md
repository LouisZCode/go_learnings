# listapp-evolution

The list app built up in small, runnable steps. Each step is a folder containing a
runnable Go program. Diff any two steps to see exactly what the new concept added.

| Step | Adds | Status |
|------|------|--------|
| `00_basic_http` | Go backend answering HTTP: router, handlers, JSON in/out, list kept in memory | ✅ you are here |
| `01_crud_complete` | update (PUT), validation, consistent JSON errors | next |
| `02_postgres` | the in-memory slice replaced by a real database (pgx driver) | |
| `03_frontend` | React talks to the API (→ becomes `listapp`) | |

Run a step from the repo root:

```bash
go run ./listapp-evolution/00_basic_http
```

Step 0 listens on **:8081** (listapp uses :8080), so you can run both at the same time and compare.
