package stores

import (
	"errors"
	"fmt"
	"net/http"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
	"dvdrental-api/internal/utils"

	"github.com/jackc/pgx/v5"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) FetchStores(w http.ResponseWriter, r *http.Request) {
	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchStoresParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	stores, total, err := h.service.FetchStores(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	totalPages := (int(total) + limit - 1) / limit

	utils.RespondWithJSON(w, http.StatusOK, StoresResponse{
		Stores: stores,
		PaginatedResponse: types.PaginatedResponse{
			Pagination: types.Pagination{
				Page:        page,
				Limit:       limit,
				Total:       total,
				TotalPages:  totalPages,
				HasNextPage: page < totalPages,
				HasPrevPage: page > 1,
			},
		},
		MessageResponse: types.MessageResponse{
			Message: "Stores have been fetched successfully.",
		},
	})
}

func (h *handler) GetStore(w http.ResponseWriter, r *http.Request) {
	storeID, err := utils.GetUrlID(r, "storeID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse store id: %v", err))
		return
	}

	store, err := h.service.GetStore(r.Context(), storeID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Store not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, StoreResponse{
		Store: store,
		MessageResponse: types.MessageResponse{
			Message: "Store detail has been fetched successfully.",
		},
	})
}

func (h *handler) CreateStore(w http.ResponseWriter, r *http.Request) {
	req := StoreRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.CreateStoreParams{
		ManagerStaffID: req.ManagerStaffID,
		AddressID:      req.AddressID,
	}

	store, err := h.service.CreateStore(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create store.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, StoreResponse{
		Store: store,
		MessageResponse: types.MessageResponse{
			Message: "Store has been created successfully.",
		},
	})
}

func (h *handler) UpdateStore(w http.ResponseWriter, r *http.Request) {
	storeID, err := utils.GetUrlID(r, "storeID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse store id: %v", err))
		return
	}

	req := StoreRequest{}
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateStoreParams{
		StoreID:        storeID,
		ManagerStaffID: req.ManagerStaffID,
		AddressID:      req.AddressID,
	}

	store, err := h.service.UpdateStore(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update store.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, StoreResponse{
		Store: store,
		MessageResponse: types.MessageResponse{
			Message: "Store has been updated successfully.",
		},
	})
}

func (h *handler) UpdateStorePartial(w http.ResponseWriter, r *http.Request) {
	storeID, err := utils.GetUrlID(r, "storeID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse store id: %v", err))
		return
	}

	req := UpdateStorePartialRequest{}
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateStorePartialParams{
		StoreID:        storeID,
		AddressID:      utils.ToInt2(req.AddressID),
		ManagerStaffID: utils.ToInt2(req.ManagerStaffID),
	}

	store, err := h.service.UpdateStorePartial(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update store.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, StoreResponse{
		Store: store,
		MessageResponse: types.MessageResponse{
			Message: "Store has been updated successfully.",
		},
	})
}

func (h *handler) DeleteStore(w http.ResponseWriter, r *http.Request) {
	storeID, err := utils.GetUrlID(r, "storeID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse store id: %v", err))
		return
	}

	err = h.service.DeleteStore(r.Context(), storeID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete store.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Store has been deleted successfully.",
	})
}

func (h *handler) FetchStoreInventory(w http.ResponseWriter, r *http.Request) {
	storeID, err := utils.GetUrlID(r, "storeID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse id: %v", err))
		return
	}

	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchStoreInventoryParams{
		StoreID: int16(storeID),
		Limit:   int32(limit),
		Offset:  int32(offset),
	}

	inventory, err := h.service.FetchStoreInventory(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := StoreInventoryResponse{
		Inventory: inventory,
		MessageResponse: types.MessageResponse{
			Message: "Inventory has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) FetchStoreCustomers(w http.ResponseWriter, r *http.Request) {
	storeID, err := utils.GetUrlID(r, "storeID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse id: %v", err))
		return
	}

	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchStoreCustomersParams{
		StoreID: int16(storeID),
		Limit:   int32(limit),
		Offset:  int32(offset),
	}

	customers, err := h.service.FetchStoreCustomers(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := StoreCustomersResponse{
		Customers: customers,
		MessageResponse: types.MessageResponse{
			Message: "Customers has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}
