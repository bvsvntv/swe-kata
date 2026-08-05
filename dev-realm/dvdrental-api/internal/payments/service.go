package payments

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchPayments(ctx context.Context, arg repo.FetchPaymentsParams) ([]repo.Payment, int64, error)
	CreatePayment(ctx context.Context, arg repo.CreatePaymentParams) (repo.Payment, error)
	GetPayment(ctx context.Context, paymentID int32) (repo.Payment, error)
	UpdatePayment(ctx context.Context, arg repo.UpdatePaymentParams) (repo.Payment, error)
	DeletePayment(ctx context.Context, paymentID int32) error
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) *svc {
	return &svc{
		repo: repo,
	}
}

func (s *svc) FetchPayments(ctx context.Context, arg repo.FetchPaymentsParams) ([]repo.Payment, int64, error) {
	stores, err := s.repo.FetchPayments(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountPayments(ctx)
	if err != nil {
		return nil, 0, err
	}

	return stores, total, nil
}

func (s *svc) CreatePayment(ctx context.Context, arg repo.CreatePaymentParams) (repo.Payment, error) {
	return s.repo.CreatePayment(ctx, arg)
}

func (s *svc) GetPayment(ctx context.Context, paymentID int32) (repo.Payment, error) {
	return s.repo.GetPayment(ctx, paymentID)
}

func (s *svc) UpdatePayment(ctx context.Context, arg repo.UpdatePaymentParams) (repo.Payment, error) {
	return s.repo.UpdatePayment(ctx, arg)
}

func (s *svc) DeletePayment(ctx context.Context, paymentID int32) error {
	return s.repo.DeletePayment(ctx, paymentID)
}
