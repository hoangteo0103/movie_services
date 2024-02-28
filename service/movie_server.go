package service

import (
	"context"
	"movie_services/api"
	"movie_services/store"

	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:generate mockery --name IMovieServer
type IMovieServer interface {
	GetMovie(ctx context.Context, req *api.GetMovieRequest) (*api.Movie, error)
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

func (s *MovieServer) GetMovie(ctx context.Context, req *api.GetMovieRequest) (*api.Movie, error) {
	id := req.GetId()
	movie, err := s.querier.GetMovie(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &api.Movie{
		Id:        movie.ID,
		Title:     movie.Title,
		Year:      movie.Year,
		Genres:    movie.Genres,
		CreatedAt: timestamppb.New(movie.CreatedAt),
		UpdatedAt: timestamppb.New(movie.UpdatedAt),
	}

	return resp, nil
}
