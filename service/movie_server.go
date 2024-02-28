package service

import (
	"context"
	"movie_services/api"
	"movie_services/store"

	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name IMovieServer
type IMovieServer interface {
	GetStatus(ctx context.Context, req *api.GetMovieRequest) (*api.Movie, error)
}

type MovieServer struct {
	api.UnimplementedMovieServiceServer
	db      *sqlx.DB
	querier store.Querier
}

type NewMovieServerOpt struct {
	Db      *sqlx.DB
	Querier store.Querier
}

func NewMovieServer(opt *NewMovieServerOpt) *MovieServer {
	return &MovieServer{
		db:      opt.Db,
		querier: opt.Querier,
	}
}

func (s *MovieServer) GetStatus(ctx context.Context, req *api.GetMovieRequest) (*api.Movie, error) {
	id := req.GetId()
	movie, err := s.querier.GetMovie(ctx, id)
	if err != nil {
		return nil, err
	}

	resp:= &api.Movie{
		Id: movie.ID,
		Title: movie.Title,
		Year: movie.Year,
		Genres: movie.Genres,
		create
}
