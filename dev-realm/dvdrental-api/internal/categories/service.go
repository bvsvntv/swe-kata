package categories

import (
	"context"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
)

type Service interface {
	FetchCategories(ctx context.Context, arg repo.FetchCategoriesParams) ([]repo.Category, int64, error)
	GetCategory(ctx context.Context, categoryID int32) (repo.Category, error)
	CreateCategory(ctx context.Context, name string) (repo.Category, error)
	DeleteCategory(ctx context.Context, categoryID int32) error
	UpdateCategory(ctx context.Context, arg repo.UpdateCategoryParams) (repo.Category, error)
	UpdateCategoryPartial(ctx context.Context, arg repo.UpdateCategoryPartialParams) (repo.Category, error)
	FetchCategoryFilms(ctx context.Context, arg repo.FetchCategoryFilmsParams) ([]repo.Film, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(repo repo.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) FetchCategories(ctx context.Context, arg repo.FetchCategoriesParams) ([]repo.Category, int64, error) {
	categories, err := s.repo.FetchCategories(ctx, arg)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountCategories(ctx)
	if err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (s *svc) GetCategory(ctx context.Context, categoryID int32) (repo.Category, error) {
	return s.repo.GetCategory(ctx, categoryID)
}

func (s *svc) CreateCategory(ctx context.Context, name string) (repo.Category, error) {
	return s.repo.CreateCategory(ctx, name)
}

func (s *svc) DeleteCategory(ctx context.Context, categoryID int32) error {
	return s.repo.DeleteCategory(ctx, categoryID)
}

func (s *svc) UpdateCategory(ctx context.Context, arg repo.UpdateCategoryParams) (repo.Category, error) {
	return s.repo.UpdateCategory(ctx, arg)
}

func (s *svc) UpdateCategoryPartial(ctx context.Context, arg repo.UpdateCategoryPartialParams) (repo.Category, error) {
	return s.repo.UpdateCategoryPartial(ctx, arg)
}

func (s *svc) FetchCategoryFilms(ctx context.Context, arg repo.FetchCategoryFilmsParams) ([]repo.Film, error) {
	return s.repo.FetchCategoryFilms(ctx, arg)
}
