package categories

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type CategoriesResponse struct {
	types.MessageResponse
	Categories []repo.Category `json:"categories"`
	types.PaginatedResponse
}

type CategoryResponse struct {
	types.MessageResponse
	Category repo.Category `json:"category"`
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	Name string `json:"name"`
}

type UpdateCategoryPartialRequest struct {
	Name *string `json:"name,omitempty"`
}

type CategoryFilmsResponse struct {
	types.MessageResponse
	Films []repo.Film `json:"films"`
}
