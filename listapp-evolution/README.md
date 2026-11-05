# listapp-evolution

The list app built up in small, runnable steps. Each step is a folder containing a
runnable Go program. Diff any two steps to see exactly what the new concept added.

| Step | Adds | Status |
|------|------|--------|
| `00_basic_http` | Go backend answering HTTP: router, handlers, JSON in/out, list kept in memory | ✅ done |
| `02_postgres` | the in-memory slice replaced by a real database (pgx driver + pool + SQL) | ✅ **you are here** |
| `01_crud_complete` | update (PUT), validation, consistent JSON errors | next |
| `03_frontend` | React talks to the API (→ becomes `listapp`) | |

Steps don't have to be built in number order — `01` slots in whenever we want PUT.

**Run a step** (port differs per step, so they can run side by side):

```bash
# 00 — std-lib only, belongs to the repo's learning module → runs from the root:
go run ./listapp-evolution/00_basic_http            # :8081

# 02 — has its own go.mod (it adds the pgx dependency; keeps the learning module
# dependency-free, same pattern as listapp/) → runs from inside its folder:
cd listapp-evolution/02_postgres && go run .       # :8082
```
