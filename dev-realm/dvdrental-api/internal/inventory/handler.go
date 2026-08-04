package inventory

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

func (h *handler) FetchInventories(w http.ResponseWriter, r *http.Request) {
	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchInventoriesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	inventories, total, err := h.service.FetchInventories(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := InventoriesResponse{
		Inventory: inventories,
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
			Message: "Inventories has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetInventory(w http.ResponseWriter, r *http.Request) {
	inventoryID, err := utils.GetUrlID(r, "inventoryID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse inventory id: %v", err))
		return
	}

	inventory, err := h.service.GetInventory(r.Context(), inventoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Inventory not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, InventoryResponse{
		Inventory: inventory,
		MessageResponse: types.MessageResponse{
			Message: "Inventory detail has been fetched successfully.",
		},
	})
}

func (h *handler) CreateInventory(w http.ResponseWriter, r *http.Request) {
	req := InventoryRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.CreateInventoryParams{
		FilmID:  req.FilmID,
		StoreID: req.StoreID,
	}

	inventory, err := h.service.CreateInventory(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create inventory.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, InventoryResponse{
		Inventory: inventory,
		MessageResponse: types.MessageResponse{
			Message: "Inventory has been created successfully.",
		},
	})
}

func (h *handler) DeleteInventory(w http.ResponseWriter, r *http.Request) {
	inventoryID, err := utils.GetUrlID(r, "inventoryID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse inventory id: %v", err))
		return
	}

	err = h.service.DeleteInventory(r.Context(), inventoryID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete inventory.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Inventory has been deleted successfully.",
	})
}
