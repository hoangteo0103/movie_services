// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package store

import (
	"context"

	"github.com/lib/pq"
)

const DeleteMovie = `-- name: DeleteMovie :exec
UPDATE movies
SET deleted_at = now()
WHERE id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteMovieStmt, DeleteMovie, id)
	return err
}

const GetMovie = `-- name: GetMovie :one
SELECT id, title, year, runtime, genres, created_at, updated_at, deleted_at FROM movies 
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMovie(ctx context.Context, id int32) (Movie, error) {
	row := q.queryRow(ctx, q.getMovieStmt, GetMovie, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Year,
		&i.Runtime,
		pq.Array(&i.Genres),
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const ListMovies = `-- name: ListMovies :many
SELECT id, title, year, runtime, genres, created_at, updated_at, deleted_at FROM movies 
WHERE genres && $3 
ORDER BY runtime
LIMIT $1 OFFSET $2
`

type ListMoviesParams struct {
	Limit  int64    `db:"limit" json:"limit"`
	Offset int64    `db:"offset" json:"offset"`
	Genres []string `db:"genres" json:"genres"`
}

func (q *Queries) ListMovies(ctx context.Context, arg ListMoviesParams) ([]Movie, error) {
	rows, err := q.query(ctx, q.listMoviesStmt, ListMovies, arg.Limit, arg.Offset, pq.Array(arg.Genres))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Year,
			&i.Runtime,
			pq.Array(&i.Genres),
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateMovie = `-- name: UpdateMovie :exec
UPDATE movies 
SET title = CASE WHEN $1::boolean THEN $2::VARCHAR(50) ELSE title END,
    year = CASE WHEN $3::boolean THEN $4::INTEGER ELSE year END,
    runtime = CASE WHEN $5::boolean THEN $6::INTEGER ELSE runtime END
WHERE id = $7
RETURNING id, title, year, runtime, genres, created_at, updated_at, deleted_at
`

type UpdateMovieParams struct {
	UpdateTitle   bool   `db:"update_title" json:"update_title"`
	Title         string `db:"title" json:"title"`
	UpdateYear    bool   `db:"update_year" json:"update_year"`
	Year          int32  `db:"year" json:"year"`
	UpdateRuntime bool   `db:"update_runtime" json:"update_runtime"`
	Runtime       int32  `db:"runtime" json:"runtime"`
	ID            int32  `db:"id" json:"id"`
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) error {
	_, err := q.exec(ctx, q.updateMovieStmt, UpdateMovie,
		arg.UpdateTitle,
		arg.Title,
		arg.UpdateYear,
		arg.Year,
		arg.UpdateRuntime,
		arg.Runtime,
		arg.ID,
	)
	return err
}