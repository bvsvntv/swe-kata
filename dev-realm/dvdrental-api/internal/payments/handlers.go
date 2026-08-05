package payments

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

func (h *handler) FetchPayments(w http.ResponseWriter, r *http.Request) {
	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchPaymentsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	payments, total, err := h.service.FetchPayments(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	totalPages := (int(total) + limit - 1) / limit

	utils.RespondWithJSON(w, http.StatusOK, PaymentsResponse{
		Payments: payments,
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
			Message: "Payments have been fetched successfully.",
		},
	})
}

func (h *handler) GetPayment(w http.ResponseWriter, r *http.Request) {
	paymentID, err := utils.GetUrlID(r, "paymentID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse payment id: %v", err))
		return
	}

	payment, err := h.service.GetPayment(r.Context(), paymentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Payment not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, PaymentResponse{
		Payment: payment,
		MessageResponse: types.MessageResponse{
			Message: "Payment detail has been fetched successfully.",
		},
	})
}

func (h *handler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	req := PaymentRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.CreatePaymentParams{
		CustomerID: req.CustomerID,
		StaffID:    req.StaffID,
		RentalID:   req.RentalID,
		Amount:     req.Amount,
	}

	payment, err := h.service.CreatePayment(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create payment.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, PaymentResponse{
		Payment: payment,
		MessageResponse: types.MessageResponse{
			Message: "Payment has been created successfully.",
		},
	})
}

func (h *handler) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	paymentID, err := utils.GetUrlID(r, "paymentID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse payment id: %v", err))
		return
	}

	req := PaymentRequest{}
	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdatePaymentParams{
		PaymentID:  paymentID,
		CustomerID: req.CustomerID,
		StaffID:    req.StaffID,
		RentalID:   req.RentalID,
		Amount:     req.Amount,
	}

	payment, err := h.service.UpdatePayment(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update payment.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, PaymentResponse{
		Payment: payment,
		MessageResponse: types.MessageResponse{
			Message: "Payment has been updated successfully.",
		},
	})
}

func (h *handler) DeletePayment(w http.ResponseWriter, r *http.Request) {
	paymentID, err := utils.GetUrlID(r, "paymentID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse payment id: %v", err))
		return
	}

	err = h.service.DeletePayment(r.Context(), paymentID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete payment.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Payment has been deleted successfully.",
	})
}
