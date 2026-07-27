package actors

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchActors(ctx context.Context, arg repo.FetchActorsParams) ([]repo.Actor, int64, error)
	GetActor(ctx context.Context, actorID int32) (repo.Actor, error)
}

type svc struct {
	// repository
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) FetchActors(ctx context.Context, arg repo.FetchActorsParams) ([]repo.Actor, int64, error) {
	actors, err := s.repo.FetchActors(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountActors(ctx)
	if err != nil {
		return nil, 0, err
	}

	return actors, total, nil
}

func (s *svc) GetActor(ctx context.Context, actorID int32) (repo.Actor, error) {
	return s.repo.GetActor(ctx, actorID)
}
