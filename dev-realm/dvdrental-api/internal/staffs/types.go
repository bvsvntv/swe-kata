package staffs

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type StaffsResponse struct {
	types.MessageResponse
	Staffs []repo.Staff `json:"staffs"`
	types.PaginatedResponse
}

type StaffResponse struct {
	types.MessageResponse
	Staff repo.Staff `json:"staff"`
}

type StaffRequest struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     *string `json:"email"`
	Active    bool    `json:"active"`
	Username  string  `json:"username"`
	Password  *string `json:"password"`
	Picture   *[]byte `json:"picture"`
	AddressID int16   `json:"address_id"`
	StoreID   int16   `json:"store_id"`
}

type UpdateStaffPartialRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Active    *bool   `json:"active,omitempty"`
	Username  *string `json:"username,omitempty"`
	Password  *string `json:"password,omitempty"`
	Picture   *[]byte `json:"picture,omitempty"`
	AddressID *int16  `json:"address_id,omitempty"`
	StoreID   *int16  `json:"store_id,omitempty"`
}

type StaffRentalsResponse struct {
	types.MessageResponse
	Rentals []repo.Rental `json:"rentals"`
}
