package films

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	GetFilms(ctx context.Context) ([]repo.Film, error)
}

type svc struct {
	// repository
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) GetFilms(ctx context.Context) ([]repo.Film, error) {
	return s.repo.GetFilms(ctx)
}
