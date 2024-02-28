package service

import (
	"context"
	"database/sql"
	"fmt"
	"movie_services/api"
	"movie_services/store"

	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//go:generate mockery --name IMovieServer
type IMovieServer interface {
	GetMovie(ctx context.Context, req *api.GetMovieRequest) (*api.Movie, error)
	DeleteMovie(ctx context.Context, req *api.DeleteMovieRequest) (*api.DeleteMovieRequest, error)
	CreateMovie(ctx context.Context, req *api.CreateMovieRequest) (*api.Movie, error)
	UpdateMovie(ctx context.Context, req *api.UpdateMovieRequest) (*api.Movie, error)
	ListMovie(ctx context.Context, req *api.ListMoviesRequest) (*api.ListMoviesResponse, error)
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

func (s *MovieServer) CreateMovie(ctx context.Context, req *api.CreateMovieRequest) (*api.Movie, error) {
	movie := store.Movie{
		Title:   req.GetTitle(),
		Year:    req.GetYear(),
		Runtime: req.GetRuntime(),
		Genres:  req.GetGenres(),
	}

	movie, err := s.querier.CreateMovie(ctx, store.CreateMovieParams{
		Title:   req.GetTitle(),
		Year:    req.GetYear(),
		Runtime: req.GetRuntime(),
		Genres:  req.GetGenres(),
	})
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

func (s *MovieServer) DeleteMovie(ctx context.Context, req *api.DeleteMovieRequest) (*api.DeleteMovieResponse, error) {
	id := req.GetId()
	err := s.querier.DeleteMovie(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &api.DeleteMovieResponse{
		Message: "Movie deleted",
	}

	return resp, nil
}

func (s *MovieServer) UpdateMovie(ctx context.Context, req *api.UpdateMovieRequest) (*api.Movie, error) {
	id := req.GetId()
	movie, err := s.querier.UpdateMovie(ctx, store.UpdateMovieParams{
		ID:      id,
		Title:   sql.NullString{String: req.GetTitle(), Valid: true},
		Year:    sql.NullInt32{Int32: req.GetYear(), Valid: true},
		Runtime: sql.NullInt32{Int32: req.GetRuntime(), Valid: true},
	})

	fmt.Println(req.GetRuntime())
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

func (s *MovieServer) ListMovies(ctx context.Context, req *api.ListMoviesRequest) (*api.ListMoviesResponse, error) {

	limit := req.GetLimit()
	offset := req.GetOffset()
	filter := req.GetFilter()
	if filter == nil {
		filter = []string{}
	}

	sort := req.GetSortBy()
	movies, err := s.querier.ListMovies(ctx, store.ListMoviesParams{
		Limit:  int64(limit),
		Offset: int64(offset),
		Filter: filter,
		SortBy: sql.NullString{String: sort, Valid: true},
	})

	if err != nil {
		return nil, err
	}

	respMovies := make([]*api.Movie, 0, len(movies))
	for _, movie := range movies {
		respMovies = append(respMovies, &api.Movie{
			Id:        movie.ID,
			Title:     movie.Title,
			Year:      movie.Year,
			Genres:    movie.Genres,
			CreatedAt: timestamppb.New(movie.CreatedAt),
			UpdatedAt: timestamppb.New(movie.UpdatedAt),
		})
	}
	resp := &api.ListMoviesResponse{
		Movies: respMovies,
	}

	return resp, nil

}
