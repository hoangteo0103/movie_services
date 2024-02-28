package service

import (
	"context"
	"movie_services/api"
	"movie_services/store"

	"github.com/jmoiron/sqlx"
)

//go:generate mockery --name IHealthCheckServer
type IHealthCheckServer interface {
	GetStatus(ctx context.Context, req *api.HealthCheckRequest) (*api.HealthCheckStatus, error)
}

type HealthCheckServer struct {
	api.UnimplementedHealthCheckServiceServer
	db      *sqlx.DB
	querier store.Querier
}

type NewHealthCheckServerOpt struct {
	Db      *sqlx.DB
	Querier store.Querier
}

func NewHealthCheckServer(opt *NewHealthCheckServerOpt) *HealthCheckServer {
	return &HealthCheckServer{
		db:      opt.Db,
		querier: opt.Querier,
	}
}

func (s *HealthCheckServer) GetStatus(ctx context.Context, req *api.HealthCheckRequest) (*api.HealthCheckStatus, error) {
	return &api.HealthCheckStatus{
		Status:      "available",
		Environment: "dev",
		Version:     "1.0.0",
	}, nil
}
