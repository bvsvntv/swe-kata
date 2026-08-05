package payments

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"

	"github.com/jackc/pgx/v5/pgtype"
)

type PaymentsResponse struct {
	types.MessageResponse
	Payments []repo.Payment `json:"payments"`
	types.PaginatedResponse
}

type PaymentResponse struct {
	types.MessageResponse
	Payment repo.Payment `json:"payment"`
}

type PaymentRequest struct {
	CustomerID int16          `json:"customer_id"`
	RentalID   int32          `json:"rental_id"`
	StaffID    int16          `json:"staff_id"`
	Amount     pgtype.Numeric `json:"amount"`
}
