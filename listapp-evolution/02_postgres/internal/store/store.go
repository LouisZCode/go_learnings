// Package store is step 02's database layer — the only place that knows SQL exists.
//
// Compare with step 00: everything the slice did, one SQL statement does now.
package store

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrNotFound arrives when a DELETE hits no row. The handler translates
// it into a 404 — driver error → domain error → HTTP status.
var ErrNotFound = errors.New("item not found")

// Item is one element of the list.
type Item struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

// ItemStore wraps a pgxpool: a small set of connections to Postgres that
// are opened once and reused for every request. Safe for concurrent use —
// unlike step 00's global slice, which would corrupt under parallel writes.
type ItemStore struct {
	pool *pgxpool.Pool
}

// Open connects to Postgres and pings it, so a wrong address fails here,
// at startup — not on the first request.
func Open(databaseURL string) (*ItemStore, error) {
	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}
	return &ItemStore{pool: pool}, nil
}

func (s *ItemStore) Close() { s.pool.Close() }

// List returns every item, oldest first.
// Step 00 just wrote the slice out; here we BUILD a slice from the rows.
func (s *ItemStore) List(ctx context.Context) ([]Item, error) {
	rows, err := s.pool.Query(ctx, `SELECT id, title FROM items ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Item{} // never nil, so an empty table still answers []
	for rows.Next() {
		var it Item
		// Scan copies the columns of one row into Go variables.
		if err := rows.Scan(&it.ID, &it.Title); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, rows.Err()
}

// Create inserts a new row and hands back the stored item —
// RETURNING is how the caller learns the id Postgres generated.
// Step 00 did "last id + 1" by hand; the database does that job now.
func (s *ItemStore) Create(ctx context.Context, title string) (Item, error) {
	var it Item
	err := s.pool.QueryRow(ctx,
		`INSERT INTO items (title) VALUES ($1) RETURNING id, title`,
		title,
	).Scan(&it.ID, &it.Title)
	return it, err
}

// Delete removes one row. No loop needed — WHERE does the lookup.
func (s *ItemStore) Delete(ctx context.Context, id int64) error {
	tag, err := s.pool.Exec(ctx, `DELETE FROM items WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound // 0 rows deleted = the id never existed
	}
	return nil
}
