package rentals

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchRentals(ctx context.Context, arg repo.FetchRentalsParams) ([]repo.Rental, int64, error)
	GetRental(ctx context.Context, rentalID int32) (repo.Rental, error)
	CreateRental(ctx context.Context, arg repo.CreateRentalParams) (repo.Rental, error)
	UpdateRental(ctx context.Context, arg repo.UpdateRentalParams) (repo.Rental, error)
	DeleteRental(ctx context.Context, rentalID int32) error
	ReturnRental(ctx context.Context, rentalID int32) (repo.Rental, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{
		repo: repo,
	}
}

func (s *svc) FetchRentals(ctx context.Context, arg repo.FetchRentalsParams) ([]repo.Rental, int64, error) {
	rentals, err := s.repo.FetchRentals(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountRentals(ctx)
	if err != nil {
		return nil, 0, err
	}

	return rentals, total, nil
}

func (s *svc) GetRental(ctx context.Context, rentalID int32) (repo.Rental, error) {
	return s.repo.GetRental(ctx, rentalID)
}

func (s *svc) CreateRental(ctx context.Context, arg repo.CreateRentalParams) (repo.Rental, error) {
	return s.repo.CreateRental(ctx, arg)
}

func (s *svc) UpdateRental(ctx context.Context, arg repo.UpdateRentalParams) (repo.Rental, error) {
	return s.repo.UpdateRental(ctx, arg)
}

func (s *svc) DeleteRental(ctx context.Context, rentalID int32) error {
	return s.repo.DeleteRental(ctx, rentalID)
}

func (s *svc) ReturnRental(ctx context.Context, rentalID int32) (repo.Rental, error) {
	return s.repo.ReturnRental(ctx, rentalID)
}
