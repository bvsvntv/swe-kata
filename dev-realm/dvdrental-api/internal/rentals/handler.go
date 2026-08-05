package rentals

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

func (h *handler) FetchRentals(w http.ResponseWriter, r *http.Request) {
	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchRentalsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	rentals, total, err := h.service.FetchRentals(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := RentalsResponse{
		Rentals: rentals,
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
			Message: "Rentals has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetRental(w http.ResponseWriter, r *http.Request) {
	rentalID, err := utils.GetUrlID(r, "rentalID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse rental id: %v", err))
		return
	}

	rental, err := h.service.GetRental(r.Context(), int32(rentalID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Rental not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, RentalResponse{
		Rental: rental,
		MessageResponse: types.MessageResponse{
			Message: "Rental detail has been fetched successfully.",
		},
	})
}

func (h *handler) CreateRental(w http.ResponseWriter, r *http.Request) {
	req := RentalRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.CreateRentalParams{
		RentalDate:  utils.ToTimestamp(&req.RentalDate),
		InventoryID: req.InventoryID,
		CustomerID:  req.CustomerID,
		StaffID:     req.StaffID,
		ReturnDate:  utils.ToTimestamp(req.ReturnDate),
	}

	rental, err := h.service.CreateRental(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create rental.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, RentalResponse{
		Rental: rental,
		MessageResponse: types.MessageResponse{
			Message: "Rental has been created successfully.",
		},
	})
}

func (h *handler) UpdateRental(w http.ResponseWriter, r *http.Request) {
	rentalID, err := utils.GetUrlID(r, "rentalID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse rental id: %v", err))
		return
	}

	req := RentalRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateRentalParams{
		RentalID:    int32(rentalID),
		RentalDate:  utils.ToTimestamp(&req.RentalDate),
		InventoryID: req.InventoryID,
		CustomerID:  req.CustomerID,
		StaffID:     req.StaffID,
		ReturnDate:  utils.ToTimestamp(req.ReturnDate),
	}

	rental, err := h.service.UpdateRental(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update rental.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, RentalResponse{
		Rental: rental,
		MessageResponse: types.MessageResponse{
			Message: "Rental has been updated successfully.",
		},
	})
}

func (h *handler) DeleteRental(w http.ResponseWriter, r *http.Request) {
	rentalID, err := utils.GetUrlID(r, "rentalID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse rental id: %v", err))
		return
	}

	err = h.service.DeleteRental(r.Context(), int32(rentalID))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete rental.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Rental has been deleted successfully.",
	})
}

func (h *handler) ReturnRental(w http.ResponseWriter, r *http.Request) {
	rentalID, err := utils.GetUrlID(r, "rentalID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse rental id: %v", err))
		return
	}

	rental, err := h.service.ReturnRental(r.Context(), int32(rentalID))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to return rental.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, RentalResponse{
		Rental: rental,
		MessageResponse: types.MessageResponse{
			Message: "Rental has been returned successfully.",
		},
	})
}
