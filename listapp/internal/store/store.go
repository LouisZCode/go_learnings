// Package store provides database access for the list app's items.
// It is the only place that knows SQL exists.
package store

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ErrNotFound is returned by Update and Delete when the id does not exist.
var ErrNotFound = errors.New("item not found")

// Item is one element of the list. The json tags define the API shape.
type Item struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// ItemStore wraps the connection pool. It is safe for concurrent use.
type ItemStore struct {
	pool *pgxpool.Pool
}

// Open connects to Postgres and verifies the connection with a ping.
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
func (s *ItemStore) List(ctx context.Context) ([]Item, error) {
	rows, err := s.pool.Query(ctx, `SELECT id, title, done FROM items ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Item{} // never nil, so the API always renders "[]"
	for rows.Next() {
		var it Item
		if err := rows.Scan(&it.ID, &it.Title, &it.Done); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, rows.Err()
}

// Create inserts a new item and returns it with its database-generated id.
func (s *ItemStore) Create(ctx context.Context, title string) (Item, error) {
	var it Item
	err := s.pool.QueryRow(ctx,
		`INSERT INTO items (title) VALUES ($1) RETURNING id, title, done`,
		title,
	).Scan(&it.ID, &it.Title, &it.Done)
	return it, err
}

// Update replaces title and done on the given item and returns the new state.
func (s *ItemStore) Update(ctx context.Context, id int64, title string, done bool) (Item, error) {
	var it Item
	err := s.pool.QueryRow(ctx,
		`UPDATE items SET title = $1, done = $2 WHERE id = $3 RETURNING id, title, done`,
		title, done, id,
	).Scan(&it.ID, &it.Title, &it.Done)
	if errors.Is(err, pgx.ErrNoRows) {
		return Item{}, ErrNotFound
	}
	return it, err
}

// Delete removes the item. ErrNotFound if no row had that id.
func (s *ItemStore) Delete(ctx context.Context, id int64) error {
	tag, err := s.pool.Exec(ctx, `DELETE FROM items WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
