package cities

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchCities(ctx context.Context, arg repo.FetchCitiesParams) ([]repo.City, int64, error)
	GetCity(ctx context.Context, cityID int32) (repo.City, error)
	CreateCity(ctx context.Context, arg repo.CreateCityParams) (repo.City, error)
	DeleteCity(ctx context.Context, cityID int32) error
	UpdateCity(ctx context.Context, arg repo.UpdateCityParams) (repo.City, error)
	UpdateCityPartial(ctx context.Context, arg repo.UpdateCityPartialParams) (repo.City, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) FetchCities(ctx context.Context, arg repo.FetchCitiesParams) ([]repo.City, int64, error) {
	cities, err := s.repo.FetchCities(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountCities(ctx)
	if err != nil {
		return nil, 0, err
	}

	return cities, total, err
}

func (s *svc) GetCity(ctx context.Context, cityID int32) (repo.City, error) {
	return s.repo.GetCity(ctx, cityID)
}

func (s *svc) CreateCity(ctx context.Context, arg repo.CreateCityParams) (repo.City, error) {
	return s.repo.CreateCity(ctx, arg)
}

func (s *svc) DeleteCity(ctx context.Context, cityID int32) error {
	return s.repo.DeleteCity(ctx, cityID)
}

func (s *svc) UpdateCity(ctx context.Context, arg repo.UpdateCityParams) (repo.City, error) {
	return s.repo.UpdateCity(ctx, arg)
}

func (s *svc) UpdateCityPartial(ctx context.Context, arg repo.UpdateCityPartialParams) (repo.City, error) {
	return s.repo.UpdateCityPartial(ctx, arg)
}
