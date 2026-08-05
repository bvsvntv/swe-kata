package inventory

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type InventoriesResponse struct {
	types.MessageResponse
	Inventory []repo.Inventory `json:"inventories"`
	types.PaginatedResponse
}

type InventoryResponse struct {
	types.MessageResponse
	Inventory repo.Inventory `json:"inventory"`
}

type InventoryRequest struct {
	FilmID  int16 `json:"film_id"`
	StoreID int16 `json:"store_id"`
}

type InventoryRentalsResponse struct {
	types.MessageResponse
	Rentals []repo.Rental `json:"rentals"`
}
