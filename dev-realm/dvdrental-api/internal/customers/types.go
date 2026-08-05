package customers

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type CustomersResponse struct {
	types.MessageResponse
	Customers []repo.Customer `json:"customers"`
	types.PaginatedResponse
}

type CustomerResponse struct {
	types.MessageResponse
	Customer repo.Customer `json:"customer"`
}

type CustomerRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     *string `json:"email"`
	Active    bool    `json:"active"`
	AddressID int16   `json:"address_id"`
	StoreID   int16   `json:"store_id"`
}

type UpdateCustomerPartialRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Active    *bool   `json:"active,omitempty"`
	AddressID *int16  `json:"address_id,omitempty"`
	StoreID   *int16  `json:"store_id,omitempty"`
}

type CustomerRentalsResponse struct {
	types.MessageResponse
	Rentals []repo.Rental `json:"rentals"`
}

type CustomerPaymentsResponse struct {
	types.MessageResponse
	Payments []repo.Payment `json:"payments"`
}
