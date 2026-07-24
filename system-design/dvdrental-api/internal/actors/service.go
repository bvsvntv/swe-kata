package actors

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	GetActors(ctx context.Context, arg repo.GetActorsParams) ([]repo.Actor, int64, error)
}

type svc struct {
	// repository
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) GetActors(ctx context.Context, arg repo.GetActorsParams) ([]repo.Actor, int64, error) {
	actors, err := s.repo.GetActors(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountActors(ctx)
	if err != nil {
		return nil, 0, err
	}

	return actors, total, nil
}
