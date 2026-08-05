package customers

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchCustomers(ctx context.Context, arg repo.FetchCustomersParams) ([]repo.Customer, int64, error)
	GetCustomer(ctx context.Context, customerID int32) (repo.Customer, error)
	CreateCustomer(ctx context.Context, arg repo.CreateCustomerParams) (repo.Customer, error)
	UpdateCustomer(ctx context.Context, arg repo.UpdateCustomerParams) (repo.Customer, error)
	UpdateCustomerPartial(ctx context.Context, arg repo.UpdateCustomerPartialParams) (repo.Customer, error)
	DeleteCustomer(ctx context.Context, customerID int32) error
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) *svc {
	return &svc{
		repo: repo,
	}
}

func (s *svc) FetchCustomers(ctx context.Context, arg repo.FetchCustomersParams) ([]repo.Customer, int64, error) {
	customers, err := s.repo.FetchCustomers(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountCustomers(ctx)
	if err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

func (s *svc) GetCustomer(ctx context.Context, customerID int32) (repo.Customer, error) {
	return s.repo.GetCustomer(ctx, customerID)
}

func (s *svc) CreateCustomer(ctx context.Context, arg repo.CreateCustomerParams) (repo.Customer, error) {
	return s.repo.CreateCustomer(ctx, arg)
}

func (s *svc) UpdateCustomer(ctx context.Context, arg repo.UpdateCustomerParams) (repo.Customer, error) {
	return s.repo.UpdateCustomer(ctx, arg)
}

func (s *svc) UpdateCustomerPartial(ctx context.Context, arg repo.UpdateCustomerPartialParams) (repo.Customer, error) {
	return s.repo.UpdateCustomerPartial(ctx, arg)
}

func (s *svc) DeleteCustomer(ctx context.Context, customerID int32) error {
	return s.repo.DeleteCustomer(ctx, customerID)
}
