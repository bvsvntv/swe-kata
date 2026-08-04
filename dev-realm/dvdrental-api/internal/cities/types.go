package cities

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type CitiesResponse struct {
	types.MessageResponse
	Cities []repo.City `json:"cities"`
	types.PaginatedResponse
}

type CityResponse struct {
	types.MessageResponse
	City repo.City `json:"city"`
}

type CityRequest struct {
	CountryID int16  `json:"country_id"`
	City      string `json:"city"`
}

type UpdateCityPartialRequest struct {
	CountryID *int16  `json:"country_id,omitempty"`
	City      *string `json:"city,omitempty"`
}
