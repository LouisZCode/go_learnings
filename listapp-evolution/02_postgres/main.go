// Step 02 — the step 00 API, but the in-memory slice is gone.
//
// The data now lives in a separate program (the Postgres server on :5432).
// This Go program is just one of its clients. What changed vs step 00:
//
//   - store.Open   connects at startup (fail fast if Postgres is down)
//   - handlers call the store with r.Context() so a hung-up browser
//     also cancels the running SQL query
//   - a.step 00's append / slice-cut / id-counter became
//     INSERT / DELETE WHERE / RETURNING in internal/store/store.go
//
// The router, the handler signature, writeJSON/fail — all unchanged.
// That's the architecture lesson: the storage swapped out underneath
// an identical API surface.
//
// It uses the SAME listapp database as the real app, so rows you add here
// appear in the app too (one database server, many clients):
//
//	PGPASSWORD=listapp psql -h localhost -U listapp -d listapp -c "SELECT * FROM items;"
//
// Run (from this folder):  go run .     — listens on :8082
package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"step02_postgres/internal/store"
)

func main() {
	dbURL := envOr("DATABASE_URL", "postgres://listapp:listapp@localhost:5432/listapp?sslmode=disable")

	items, err := store.Open(dbURL)
	if err != nil {
		log.Fatalf("connect to Postgres: %v (is it running on :5432?)", err)
	}
	defer items.Close()

	a := &api{store: items}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/items", a.list)
	mux.HandleFunc("POST /api/items", a.create)
	mux.HandleFunc("DELETE /api/items/{id}", a.remove)

	log.Println("step 02 listening on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", logRequests(mux)))
}

// api carries the store — handlers are methods on it.
type api struct {
	store *store.ItemStore
}

func (a *api) list(w http.ResponseWriter, r *http.Request) {
	all, err := a.store.List(r.Context())
	if err != nil {
		serverError(w, r, err)
		return
	}
	writeJSON(w, http.StatusOK, all)
}

func (a *api) create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fail(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	title := strings.TrimSpace(body.Title)
	if title == "" {
		fail(w, http.StatusBadRequest, "title is required")
		return
	}
	item, err := a.store.Create(r.Context(), title)
	if err != nil {
		serverError(w, r, err)
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func (a *api) remove(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		fail(w, http.StatusBadRequest, "id must be a positive number")
		return
	}
	err = a.store.Delete(r.Context(), id)
	if errors.Is(err, store.ErrNotFound) {
		fail(w, http.StatusNotFound, "item not found")
		return
	}
	if err != nil {
		serverError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// --- helpers (identical to step 00) -----------------------------------------

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func fail(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// serverError logs the real cause (never shown to the caller)
// and answers a generic 500.
func serverError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%s %s: %v", r.Method, r.URL.Path, err)
	fail(w, http.StatusInternalServerError, "internal server error")
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (%s)", r.Method, r.URL.Path, time.Since(start).Round(time.Millisecond))
	})
}
