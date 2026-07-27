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
