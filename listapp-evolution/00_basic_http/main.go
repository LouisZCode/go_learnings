// Step 0 — a basic Go backend that answers HTTP calls.
//
// One file, standard library only, no database: the list is a plain slice
// in memory, so it resets to the seed items every time you restart.
//
// This is Lesson B1 made runnable:
//
//	request → router (ServeMux) → handler function → response
//
// Run from the repo root:   go run ./listapp-evolution/00_basic_http
// Then, from another terminal:
//
//	curl localhost:8081/api/items
//	curl -X POST localhost:8081/api/items -H 'Content-Type: application/json' -d '{"title":"Coffee"}'
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// item is one element of the list.
// The json tags decide the field names the outside world sees.
type item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// The whole "database" of step 0: a slice in memory.
// Restart the program and it is back to these two entries.
var items = []item{
	{ID: 1, Title: "Understand step 0"},
	{ID: 2, Title: "This list resets on restart"},
}

func main() {
	// The router. Each line registers: "requests with this METHOD + PATH go here".
	// {id} is a wildcard — a request to /api/items/7 matches it,
	// and inside the handler r.PathValue("id") returns "7".
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/items", list)
	mux.HandleFunc("POST /api/items", create)
	mux.HandleFunc("DELETE /api/items/{id}", remove)

	// logRequests wraps the mux (see bottom): every request passes through it.
	// :8081 so this can run next to listapp (:8080).
	log.Println("step 0 listening on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", logRequests(mux)))
}

// --- handlers: read from r (request), write to w (response) -----------------

// GET /api/items — return the whole list as JSON.
func list(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, items)
}

// POST /api/items — read {"title": "..."} from the body, add a new item.
func create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string `json:"title"`
	}
	// Decode turns the JSON bytes from the request body into the struct.
	// A body that isn't valid JSON ends the request right here with a 400.
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fail(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if body.Title == "" {
		fail(w, http.StatusBadRequest, "title is required")
		return
	}

	// New id = last item's id + 1. (0 if the list is empty.)
	next := 1
	if len(items) > 0 {
		next = items[len(items)-1].ID + 1
	}
	newItem := item{ID: next, Title: body.Title}
	items = append(items, newItem)

	// 201 Created, and the caller gets the stored item back —
	// that's how the caller learns the generated id.
	writeJSON(w, http.StatusCreated, newItem)
}

// DELETE /api/items/{id} — remove one item.
func remove(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id")) // "7" → 7
	if err != nil {
		fail(w, http.StatusBadRequest, "id must be a number")
		return
	}

	for i, it := range items {
		if it.ID == id {
			// The slice trick to delete element i: glue the part
			// before i to the part after i. items[:i] + items[i+1:].
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // 204: success, on purpose no body
			return
		}
	}
	fail(w, http.StatusNotFound, "item not found")
}

// --- small helpers ----------------------------------------------------------

// writeJSON is the single place that knows how to answer with JSON:
// set the header, set the status code, encode the value.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// fail answers with {"error": "..."} and the given status.
func fail(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// logRequests is middleware: it wraps another handler and runs
// one line of logging around it. That's the whole middleware idea.
func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path)
	})
}
