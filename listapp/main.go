// Command listapp serves a small CRUD list API backed by Postgres.
//
// Routes (Go 1.22+ ServeMux patterns: "METHOD /path/{wildcard}"):
//
//	GET    /api/items       list all items
//	POST   /api/items       create an item   body: {"title": "..."}
//	PUT    /api/items/{id}  update an item   body: {"title": "...", "done": true}
//	DELETE /api/items/{id}  delete an item
//	GET    /                the built frontend (web/dist), if present
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

	"listapp/internal/store"
)

func main() {
	addr := ":" + envOr("PORT", "8080")
	dbURL := envOr("DATABASE_URL", "postgres://listapp:listapp@localhost:5432/listapp?sslmode=disable")

	items, err := store.Open(dbURL)
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}
	defer items.Close()

	api := &api{store: items}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/items", api.list)
	mux.HandleFunc("POST /api/items", api.create)
	mux.HandleFunc("PUT /api/items/{id}", api.update)
	mux.HandleFunc("DELETE /api/items/{id}", api.remove)
	mux.Handle("GET /", staticHandler())

	srv := &http.Server{
		Addr:              addr,
		Handler:           logRequests(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("listapp listening on http://localhost%s", addr)
	log.Fatal(srv.ListenAndServe())
}

// api carries the shared dependencies of the HTTP handlers.
type api struct {
	store *store.ItemStore
}

func (a *api) list(w http.ResponseWriter, r *http.Request) {
	all, err := a.store.List(r.Context())
	if err != nil {
		serverError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, all)
}

func (a *api) create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string `json:"title"`
	}
	if err := decodeBody(w, r, &body); err != nil {
		return // error response already written
	}
	title := strings.TrimSpace(body.Title)
	if title == "" {
		clientError(w, "title is required")
		return
	}
	item, err := a.store.Create(r.Context(), title)
	if err != nil {
		serverError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, item)
}

func (a *api) update(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}
	var body struct {
		Title string `json:"title"`
		Done  bool   `json:"done"`
	}
	if err := decodeBody(w, r, &body); err != nil {
		return
	}
	title := strings.TrimSpace(body.Title)
	if title == "" {
		clientError(w, "title is required")
		return
	}
	item, err := a.store.Update(r.Context(), id, title, body.Done)
	if errors.Is(err, store.ErrNotFound) {
		notFound(w)
		return
	}
	if err != nil {
		serverError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, item)
}

func (a *api) remove(w http.ResponseWriter, r *http.Request) {
	id, ok := pathID(w, r)
	if !ok {
		return
	}
	err := a.store.Delete(r.Context(), id)
	if errors.Is(err, store.ErrNotFound) {
		notFound(w)
		return
	}
	if err != nil {
		serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// --- helpers ---------------------------------------------------------------

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("write json: %v", err)
	}
}

// decodeBody parses the request body into dst and writes a 400 on bad input.
func decodeBody(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB cap
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dst); err != nil {
		clientError(w, "invalid JSON body: "+err.Error())
		return err
	}
	return nil
}

// pathID reads the {id} wildcard from the URL and writes a 400 if it is not a number.
func pathID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		clientError(w, "id must be a positive integer")
		return 0, false
	}
	return id, true
}

func clientError(w http.ResponseWriter, msg string) {
	writeJSON(w, http.StatusBadRequest, map[string]string{"error": msg})
}

func serverError(w http.ResponseWriter, err error) {
	log.Printf("internal error: %v", err)
	writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
}

func notFound(w http.ResponseWriter) {
	writeJSON(w, http.StatusNotFound, map[string]string{"error": "item not found"})
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// logRequests prints one line per request: method, path, duration.
func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (%s)", r.Method, r.URL.Path, time.Since(start).Round(time.Millisecond))
	})
}

// staticHandler serves the built frontend from web/dist with an SPA fallback:
// unknown paths return index.html so client-side routes keep working.
func staticHandler() http.Handler {
	fileServer := http.FileServer(http.Dir("web/dist"))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		index := "web/dist/index.html"
		if _, err := os.Stat(index); err != nil {
			http.Error(w, "frontend not built — run: cd web && npm install && npm run build (or use the Vite dev server on :5173)", http.StatusNotFound)
			return
		}
		if _, err := os.Stat("web/dist" + r.URL.Path); err != nil {
			http.ServeFile(w, r, index) // SPA fallback
			return
		}
		fileServer.ServeHTTP(w, r)
	})
}
