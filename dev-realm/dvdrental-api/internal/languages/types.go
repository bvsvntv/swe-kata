package languages

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type LanguagesResponse struct {
	types.MessageResponse
	Languages []repo.Language `json:"languages"`
	types.PaginatedResponse
}

type LanguageResponse struct {
	types.MessageResponse
	Language repo.Language `json:"language"`
}
