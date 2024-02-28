// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package store

import (
	"context"
)

type Querier interface {
	CreateMovie(ctx context.Context, arg CreateMovieParams) error
	DeleteMovie(ctx context.Context, id int32) error
	GetMovie(ctx context.Context, id int32) (Movie, error)
	ListMovies(ctx context.Context, arg ListMoviesParams) ([]Movie, error)
	UpdateMovie(ctx context.Context, arg UpdateMovieParams) error
}

var _ Querier = (*Queries)(nil)
