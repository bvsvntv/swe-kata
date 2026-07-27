package films

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type FilmsResponse struct {
	Films      []repo.Film      `json:"films"`
	Pagination types.Pagination `json:"pagination"`
}

type FilmResponse struct {
	Film repo.Film `json:"film"`
}
