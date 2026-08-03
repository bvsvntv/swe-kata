package stores

import (
	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
)

type StoresResponse struct {
	types.MessageResponse
	Stores []repo.Store `json:"stores"`
	types.PaginatedResponse
}

type StoreResponse struct {
	types.MessageResponse
	Store repo.Store `json:"store"`
}

type StoreRequest struct {
	ManagerStaffID int16 `json:"manager_staff_id"`
	AddressID      int16 `json:"address_id"`
}

type UpdateStorePartialRequest struct {
	ManagerStaffID *int16 `json:"manager_staff_id"`
	AddressID      *int16 `json:"address_id"`
}
