package films

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchFilms(ctx context.Context, arg repo.FetchFilmsParams) ([]repo.Film, int64, error)
}

type svc struct {
	// repository
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) FetchFilms(ctx context.Context, arg repo.FetchFilmsParams) ([]repo.Film, int64, error) {
	films, err := s.repo.FetchFilms(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountFilms(ctx)
	if err != nil {
		return nil, 0, err
	}

	return films, total, nil
}
