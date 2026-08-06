package films

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchFilms(ctx context.Context, arg repo.FetchFilmsParams) ([]repo.Film, int64, error)
	GetFilm(ctx context.Context, filmID int32) (repo.Film, error)
	FetchFilmActors(ctx context.Context, arg repo.FetchFilmActorsParams) ([]repo.Actor, error)
	FetchFilmCategories(ctx context.Context, arg repo.FetchFilmCategoriesParams) ([]repo.Category, error)
	FetchFilmInventory(ctx context.Context, arg repo.FetchFilmInventoryParams) ([]repo.Inventory, error)
	CreateFilm(ctx context.Context, arg repo.CreateFilmParams) (repo.Film, error)
	UpdateFilm(ctx context.Context, arg repo.UpdateFilmParams) (repo.Film, error)
	DeleteFilm(ctx context.Context, filmID int32) error
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

func (s *svc) GetFilm(ctx context.Context, filmID int32) (repo.Film, error) {
	return s.repo.GetFilm(ctx, filmID)
}

func (s *svc) FetchFilmActors(ctx context.Context, arg repo.FetchFilmActorsParams) ([]repo.Actor, error) {
	return s.repo.FetchFilmActors(ctx, arg)
}

func (s *svc) FetchFilmCategories(ctx context.Context, arg repo.FetchFilmCategoriesParams) ([]repo.Category, error) {
	return s.repo.FetchFilmCategories(ctx, arg)
}

func (s *svc) FetchFilmInventory(ctx context.Context, arg repo.FetchFilmInventoryParams) ([]repo.Inventory, error) {
	return s.repo.FetchFilmInventory(ctx, arg)
}

func (s *svc) CreateFilm(ctx context.Context, arg repo.CreateFilmParams) (repo.Film, error) {
	return s.repo.CreateFilm(ctx, arg)
}

func (s *svc) UpdateFilm(ctx context.Context, arg repo.UpdateFilmParams) (repo.Film, error) {
	return s.repo.UpdateFilm(ctx, arg)
}

func (s *svc) DeleteFilm(ctx context.Context, filmID int32) error {
	return s.repo.DeleteFilm(ctx, filmID)
}
