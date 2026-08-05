package customers

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

func (h *handler) FetchCustomers(w http.ResponseWriter, r *http.Request) {
	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchCustomersParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	customers, total, err := h.service.FetchCustomers(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	totalPages := (int(total) + limit - 1) / limit

	resp := CustomersResponse{
		Customers: customers,
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
			Message: "Customers has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (h *handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, err := utils.GetUrlID(r, "customerID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse customer id: %v", err))
		return
	}

	customer, err := h.service.GetCustomer(r.Context(), customerID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Customer not found.")
			return
		}

		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, CustomerResponse{
		Customer: customer,
		MessageResponse: types.MessageResponse{
			Message: "Customer detail has been fetched successfully.",
		},
	})
}

func (h *handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	req := CustomerRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.CreateCustomerParams{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Activebool: req.Active,
		AddressID:  req.AddressID,
		StoreID:    req.StoreID,
		Email:      utils.ToText(req.Email),
	}

	customer, err := h.service.CreateCustomer(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create customer.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, CustomerResponse{
		Customer: customer,
		MessageResponse: types.MessageResponse{
			Message: "Customer has been created successfully.",
		},
	})
}

func (h *handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, err := utils.GetUrlID(r, "customerID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse customer id: %v", err))
		return
	}

	err = h.service.DeleteCustomer(r.Context(), customerID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to delete customer.")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, types.MessageResponse{
		Message: "Customer has been deleted successfully.",
	})
}

func (h *handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, err := utils.GetUrlID(r, "customerID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse customer id: %v", err))
		return
	}

	req := CustomerRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateCustomerParams{
		CustomerID: customerID,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Activebool: req.Active,
		AddressID:  req.AddressID,
		StoreID:    req.StoreID,
		Email:      utils.ToText(req.Email),
	}

	customer, err := h.service.UpdateCustomer(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update customer.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, CustomerResponse{
		Customer: customer,
		MessageResponse: types.MessageResponse{
			Message: "Customer has been updated successfully.",
		},
	})
}

func (h *handler) UpdateCustomerPartial(w http.ResponseWriter, r *http.Request) {
	customerID, err := utils.GetUrlID(r, "customerID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse customer id: %v", err))
		return
	}

	req := UpdateCustomerPartialRequest{}

	if err := utils.DecodeJSON(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON.\nERROR: %v", err))
		return
	}

	arg := repo.UpdateCustomerPartialParams{
		CustomerID: customerID,
		FirstName:  utils.ToText(req.FirstName),
		LastName:   utils.ToText(req.LastName),
		Email:      utils.ToText(req.Email),
		AddressID:  utils.ToInt2(req.AddressID),
		StoreID:    utils.ToInt2(req.StoreID),
		Activebool: utils.ToBool(req.Active),
	}

	customer, err := h.service.UpdateCustomerPartial(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to update customer.\nERROR: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, CustomerResponse{
		Customer: customer,
		MessageResponse: types.MessageResponse{
			Message: "Customer has been updated successfully.",
		},
	})
}

func (h *handler) FetchCustomerRentals(w http.ResponseWriter, r *http.Request) {
	customerID, err := utils.GetUrlID(r, "customerID")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to parse id: %v", err))
		return
	}

	page := utils.GetQueryInt(r, "page", 1)
	limit := utils.GetQueryInt(r, "limit", 10)

	offset := (page - 1) * limit

	arg := repo.FetchCustomerRentalsParams{
		CustomerID: int16(customerID),
		Limit:      int32(limit),
		Offset:     int32(offset),
	}

	rentals, err := h.service.FetchCustomerRentals(r.Context(), arg)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	resp := CustomerRentalsResponse{
		Rentals: rentals,
		MessageResponse: types.MessageResponse{
			Message: "Rentals has been fetched successfully.",
		},
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}
