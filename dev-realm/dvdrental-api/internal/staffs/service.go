package staffs

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchStaffs(ctx context.Context, arg repo.FetchStaffsParams) ([]repo.Staff, int64, error)
	GetStaff(ctx context.Context, staffID int32) (repo.Staff, error)
	CreateStaff(ctx context.Context, arg repo.CreateStaffParams) (repo.Staff, error)
	DeleteStaff(ctx context.Context, staffID int32) error
	UpdateStaff(ctx context.Context, arg repo.UpdateStaffParams) (repo.Staff, error)
	UpdateStaffPartial(ctx context.Context, arg repo.UpdateStaffPartialParams) (repo.Staff, error)
	FetchStaffRentals(ctx context.Context, arg repo.FetchStaffRentalsParams) ([]repo.Rental, error)
	FetchStaffPayments(ctx context.Context, arg repo.FetchStaffPaymentsParams) ([]repo.Payment, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) *svc {
	return &svc{
		repo: repo,
	}
}

func (s *svc) FetchStaffs(ctx context.Context, arg repo.FetchStaffsParams) ([]repo.Staff, int64, error) {
	staffs, err := s.repo.FetchStaffs(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountStaffs(ctx)
	if err != nil {
		return nil, 0, err
	}

	return staffs, total, nil
}

func (s *svc) GetStaff(ctx context.Context, staffID int32) (repo.Staff, error) {
	return s.repo.GetStaff(ctx, staffID)
}

func (s *svc) CreateStaff(ctx context.Context, arg repo.CreateStaffParams) (repo.Staff, error) {
	return s.repo.CreateStaff(ctx, arg)
}

func (s *svc) DeleteStaff(ctx context.Context, staffID int32) error {
	return s.repo.DeleteStaff(ctx, staffID)
}

func (s *svc) UpdateStaff(ctx context.Context, arg repo.UpdateStaffParams) (repo.Staff, error) {
	return s.repo.UpdateStaff(ctx, arg)
}

func (s *svc) UpdateStaffPartial(ctx context.Context, arg repo.UpdateStaffPartialParams) (repo.Staff, error) {
	return s.repo.UpdateStaffPartial(ctx, arg)
}

func (s *svc) FetchStaffRentals(ctx context.Context, arg repo.FetchStaffRentalsParams) ([]repo.Rental, error) {
	return s.repo.FetchStaffRentals(ctx, arg)
}

func (s *svc) FetchStaffPayments(ctx context.Context, arg repo.FetchStaffPaymentsParams) ([]repo.Payment, error) {
	return s.repo.FetchStaffPayments(ctx, arg)
}
