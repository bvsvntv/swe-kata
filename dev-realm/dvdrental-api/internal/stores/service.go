package stores

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchStores(ctx context.Context, arg repo.FetchStoresParams) ([]repo.Store, int64, error)
	GetStore(ctx context.Context, storeID int32) (repo.Store, error)
	CreateStore(ctx context.Context, arg repo.CreateStoreParams) (repo.Store, error)
	UpdateStore(ctx context.Context, arg repo.UpdateStoreParams) (repo.Store, error)
	UpdateStorePartial(ctx context.Context, arg repo.UpdateStorePartialParams) (repo.Store, error)
	DeleteStore(ctx context.Context, storeID int32) error
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) FetchStores(ctx context.Context, arg repo.FetchStoresParams) ([]repo.Store, int64, error) {
	stores, err := s.repo.FetchStores(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountStores(ctx)
	if err != nil {
		return nil, 0, err
	}

	return stores, total, nil
}

func (s *svc) CreateStore(ctx context.Context, arg repo.CreateStoreParams) (repo.Store, error) {
	return s.repo.CreateStore(ctx, arg)
}

func (s *svc) GetStore(ctx context.Context, storeID int32) (repo.Store, error) {
	return s.repo.GetStore(ctx, storeID)
}

func (s *svc) UpdateStore(ctx context.Context, arg repo.UpdateStoreParams) (repo.Store, error) {
	return s.repo.UpdateStore(ctx, arg)
}

func (s *svc) UpdateStorePartial(ctx context.Context, arg repo.UpdateStorePartialParams) (repo.Store, error) {
	return s.repo.UpdateStorePartial(ctx, arg)
}

func (s *svc) DeleteStore(ctx context.Context, storeID int32) error {
	return s.repo.DeleteStore(ctx, storeID)
}
