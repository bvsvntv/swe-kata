package countries

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchCountries(ctx context.Context, arg repo.FetchCountriesParams) ([]repo.Country, int64, error)
	GetCountry(ctx context.Context, countryID int32) (repo.Country, error)
	CreateCountry(ctx context.Context, country string) (repo.Country, error)
	UpdateCountry(ctx context.Context, arg repo.UpdateCountryParams) (repo.Country, error)
	DeleteCountry(ctx context.Context, countryID int32) error
	FetchCountryCities(ctx context.Context, arg repo.FetchCountryCitiesParams) ([]repo.City, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) FetchCountries(ctx context.Context, arg repo.FetchCountriesParams) ([]repo.Country, int64, error) {
	countries, err := s.repo.FetchCountries(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountCountries(ctx)
	if err != nil {
		return nil, 0, err
	}

	return countries, total, err
}

func (s *svc) GetCountry(ctx context.Context, countryID int32) (repo.Country, error) {
	return s.repo.GetCountry(ctx, countryID)
}

func (s *svc) CreateCountry(ctx context.Context, country string) (repo.Country, error) {
	return s.repo.CreateCountry(ctx, country)
}

func (s *svc) UpdateCountry(ctx context.Context, arg repo.UpdateCountryParams) (repo.Country, error) {
	return s.repo.UpdateCountry(ctx, arg)
}

func (s *svc) DeleteCountry(ctx context.Context, countryID int32) error {
	return s.repo.DeleteCountry(ctx, countryID)
}

func (s *svc) FetchCountryCities(ctx context.Context, arg repo.FetchCountryCitiesParams) ([]repo.City, error) {
	return s.repo.FetchCountryCities(ctx, arg)
}
