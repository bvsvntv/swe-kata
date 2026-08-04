package films

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type FilmsResponse struct {
	types.MessageResponse
	Films []repo.Film `json:"films"`
	types.PaginatedResponse
}

type FilmResponse struct {
	types.MessageResponse
	Film repo.Film `json:"film"`
}

type FilmActorsResponse struct {
	types.MessageResponse
	Actors []repo.Actor `json:"actors"`
}

type FilmCategoriesResponse struct {
	types.MessageResponse
	Categoryies []repo.Category `json:"categories"`
}

type FilmInventoryResponse struct {
	types.MessageResponse
	Inventory []repo.Inventory `json:"inventory"`
}
