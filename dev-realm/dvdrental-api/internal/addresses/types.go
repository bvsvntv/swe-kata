package addresses

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type AddressesResponse struct {
	types.MessageResponse
	Addresses []repo.Address `json:"addresses"`
	types.PaginatedResponse
}

type AddressResponse struct {
	types.MessageResponse
	Address repo.Address `json:"address"`
}

type AddressRequest struct {
	Address    string  `json:"address"`
	Address2   *string `json:"address2"`
	District   string  `json:"district"`
	CityID     int32   `json:"city_id"`
	PostalCode *string `json:"postal_code"`
	Phone      string  `json:"phone"`
}

type UpdateAddressPartialRequest struct {
	Address    *string `json:"address,omitempty"`
	Address2   *string `json:"address2,omitempty"`
	District   *string `json:"district,omitempty"`
	CityID     *int32  `json:"city_id,omitempty"`
	PostalCode *string `json:"postal_code,omitempty"`
	Phone      *string `json:"phone,omitempty"`
}
