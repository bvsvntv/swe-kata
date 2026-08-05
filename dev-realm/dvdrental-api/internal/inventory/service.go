package inventory

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchInventories(ctx context.Context, arg repo.FetchInventoriesParams) ([]repo.Inventory, int64, error)
	GetInventory(ctx context.Context, inventoryID int32) (repo.Inventory, error)
	CreateInventory(ctx context.Context, arg repo.CreateInventoryParams) (repo.Inventory, error)
	DeleteInventory(ctx context.Context, inventoryID int32) error
	FetchInventoryRentals(ctx context.Context, arg repo.FetchInventoryRentalsParams) ([]repo.Rental, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) *svc {
	return &svc{
		repo: repo,
	}
}

func (s *svc) FetchInventories(ctx context.Context, arg repo.FetchInventoriesParams) ([]repo.Inventory, int64, error) {
	inventory, err := s.repo.FetchInventories(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountInventories(ctx)
	if err != nil {
		return nil, 0, err
	}

	return inventory, total, nil
}

func (s *svc) GetInventory(ctx context.Context, inventoryID int32) (repo.Inventory, error) {
	return s.repo.GetInventory(ctx, inventoryID)
}

func (s *svc) CreateInventory(ctx context.Context, arg repo.CreateInventoryParams) (repo.Inventory, error) {
	return s.repo.CreateInventory(ctx, arg)
}

func (s *svc) DeleteInventory(ctx context.Context, inventoryID int32) error {
	return s.repo.DeleteInventory(ctx, inventoryID)
}

func (s *svc) FetchInventoryRentals(ctx context.Context, arg repo.FetchInventoryRentalsParams) ([]repo.Rental, error) {
	return s.repo.FetchInventoryRentals(ctx, arg)
}
