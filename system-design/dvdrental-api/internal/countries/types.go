package countries

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type CountriesResponse struct {
	types.MessageResponse
	Countries []repo.Country `json:"countries"`
	types.PaginatedResponse
}

type CountryResponse struct {
	types.MessageResponse
	Country repo.Country `json:"country"`
}

type CreateCountryRequest struct {
	Country string `json:"country"`
}

type UpdateCountryRequest struct {
	Country string `json:"country"`
}
