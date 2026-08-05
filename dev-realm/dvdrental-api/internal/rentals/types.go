package rentals

import (
	"time"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type RentalsResponse struct {
	types.MessageResponse
	Rentals []repo.Rental `json:"rentals"`
	types.PaginatedResponse
}

type RentalResponse struct {
	types.MessageResponse
	Rental repo.Rental `json:"rental"`
}

type RentalRequest struct {
	RentalDate  time.Time  `json:"rental_date"`
	InventoryID int32      `json:"inventory_id"`
	CustomerID  int16      `json:"customer_id"`
	StaffID     int16      `json:"staff_id"`
	ReturnDate  *time.Time `json:"return_date"`
}
