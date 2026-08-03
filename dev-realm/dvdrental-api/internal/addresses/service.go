package addresses

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchAddresses(ctx context.Context, arg repo.FetchAddressesParams) ([]repo.Address, int64, error)
	GetAddress(ctx context.Context, addressID int32) (repo.Address, error)
	CreateAddress(ctx context.Context, arg repo.CreateAddressParams) (repo.Address, error)
	UpdateAddress(ctx context.Context, arg repo.UpdateAddressParams) (repo.Address, error)
	UpdateAddressPartial(ctx context.Context, arg repo.UpdateAddressPartialParams) (repo.Address, error)
	DeleteAddress(ctx context.Context, addressID int32) error
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) FetchAddresses(ctx context.Context, arg repo.FetchAddressesParams) ([]repo.Address, int64, error) {
	addresses, err := s.repo.FetchAddresses(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountAddresses(ctx)
	if err != nil {
		return nil, 0, err
	}

	return addresses, total, err
}

func (s *svc) GetAddress(ctx context.Context, addressID int32) (repo.Address, error) {
	return s.repo.GetAddress(ctx, addressID)
}

func (s *svc) CreateAddress(ctx context.Context, arg repo.CreateAddressParams) (repo.Address, error) {
	return s.repo.CreateAddress(ctx, arg)
}

func (s *svc) UpdateAddress(ctx context.Context, arg repo.UpdateAddressParams) (repo.Address, error) {
	return s.repo.UpdateAddress(ctx, arg)
}

func (s *svc) UpdateAddressPartial(ctx context.Context, arg repo.UpdateAddressPartialParams) (repo.Address, error) {
	return s.repo.UpdateAddressPartial(ctx, arg)
}

func (s *svc) DeleteAddress(ctx context.Context, addressID int32) error {
	return s.repo.DeleteAddress(ctx, addressID)
}
