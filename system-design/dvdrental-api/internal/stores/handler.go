package stores

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	repo "dvdrental-api/internal/adapters/postgresql/sqlc"
	"dvdrental-api/internal/types"
	"dvdrental-api/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

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
	storeIDString := chi.URLParam(r, "storeID")
	storeID, err := strconv.Atoi(storeIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse store id: %v", err))
		return
	}

	store, err := h.service.GetStore(r.Context(), int32(storeID))
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
	decoder := json.NewDecoder(r.Body)
	args := UpdateStoreRequest{}

	err := decoder.Decode(&args)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	store, err := h.service.CreateStore(r.Context(), repo.CreateStoreParams{
		ManagerStaffID: args.ManagerStaffID,
		AddressID:      args.AddressID,
	})
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
	storeIDString := chi.URLParam(r, "storeID")
	storeID, err := strconv.Atoi(storeIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse store id: %v", err))
		return
	}

	var args UpdateStoreRequest
	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	store, err := h.service.UpdateStore(r.Context(), repo.UpdateStoreParams{
		StoreID:        int32(storeID),
		ManagerStaffID: args.ManagerStaffID,
		AddressID:      args.AddressID,
	})
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
	storeIDString := chi.URLParam(r, "storeID")
	storeID, err := strconv.Atoi(storeIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse store id: %v", err))
		return
	}

	var args UpdateStorePartialRequest
	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	partialParams := repo.UpdateStorePartialParams{
		StoreID: int32(storeID),
	}

	if args.ManagerStaffID != nil {
		partialParams.ManagerStaffID = pgtype.Int2{
			Int16: *args.ManagerStaffID,
			Valid: true,
		}
	}

	if args.AddressID != nil {
		partialParams.AddressID = pgtype.Int2{
			Int16: *args.AddressID,
			Valid: true,
		}
	}

	store, err := h.service.UpdateStorePartial(r.Context(), partialParams)
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
	storeIDString := chi.URLParam(r, "storeID")
	storeID, err := strconv.Atoi(storeIDString)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse store id: %v", err))
		return
	}

	err = h.service.DeleteStore(r.Context(), int32(storeID))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete store.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Store has been deleted successfully.",
	})
}
